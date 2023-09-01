package lang

import (
	"github.com/LiveAlone/go-util/domain"
	"github.com/LiveAlone/go-util/domain/config"
	"github.com/LiveAlone/go-util/domain/mysql"
	"github.com/LiveAlone/go-util/domain/template"
	"github.com/LiveAlone/go-util/domain/template/bo"
	"log"
	"sort"
	"strings"
)

// JavaCodeGenerator 生成java dao
type JavaCodeGenerator struct {
	ConfigLoader   *config.Loader
	JavaConfigYaml *JavaConfigYaml
	TempGen        *template.Generator
}

func NewJavaCodeGenerator(configLoader *config.Loader, templateGenerator *template.Generator) *JavaCodeGenerator {
	java := &JavaCodeGenerator{
		ConfigLoader: configLoader,
	}
	javaConfigYaml := new(JavaConfigYaml)
	err := configLoader.LoadConfigToEntity("conf/java_model.yaml", javaConfigYaml)
	if err != nil {
		log.Fatalf("yaml file read error %v path: %v", err, "conf/java_model.yaml")
	}
	java.JavaConfigYaml = javaConfigYaml
	java.TempGen = templateGenerator
	return java
}

func (g *JavaCodeGenerator) GenDaoFromTableInfo(tableInfo *mysql.TableInfo, params *CodeGenParams) (rs map[string]string, err error) {
	rs = make(map[string]string)

	// 1. jpa model dao 生成
	modelStruct := g.covertToModelStruct(tableInfo, params)
	modelContent, err := g.TempGen.GenerateTemplateByName("JpaModel", modelStruct, nil)
	if err != nil {
		return nil, err
	}
	rs[strings.ReplaceAll(params.PackageName, ".", "/")+"/"+modelStruct.BeanName+".java"] = modelContent

	daoStruct := g.convertIndexModeStruct(tableInfo, params)
	daoContent, err := g.TempGen.GenerateTemplateByName("JpaDao", daoStruct, nil)
	if err != nil {
		return nil, err
	}
	rs[strings.ReplaceAll(params.PackageName, ".", "/")+"/"+daoStruct.BeanName+"Repository.java"] = daoContent
	return
}

func (g *JavaCodeGenerator) convertIndexModeStruct(tableInfo *mysql.TableInfo, params *CodeGenParams) *bo.DaoStruct {
	rs := &bo.DaoStruct{
		PackageName: params.PackageName,
		TableName:   tableInfo.Schema.TableName,
		BeanName:    domain.ToCamelCaseFistLarge(tableInfo.Schema.TableName),
		Comment:     tableInfo.Schema.TableComment,
	}

	// 构建字段名映射关系
	columnMap := make(map[string]*mysql.TableInfoColumn)
	for _, column := range tableInfo.Column {
		columnMap[column.ColumnName] = column
	}

	indexMap := make(map[string]*bo.DaoIndex)
	for _, statistics := range tableInfo.Index {
		current, ok := indexMap[statistics.IndexName]
		if !ok {
			current = &bo.DaoIndex{
				IndexName:    statistics.IndexName,
				Unique:       !statistics.NoUnique,
				Fields:       make([]*bo.DaoIndexField, 0),
				IndexComment: statistics.IndexComment,
			}
			indexMap[statistics.IndexName] = current
		}

		current.Fields = append(current.Fields, &bo.DaoIndexField{
			Index:       statistics.SeqInIndex,
			ColumnName:  statistics.ColumnName,
			ColumnType:  columnMap[statistics.ColumnName].DataType,
			FieldName:   domain.ToCamelCaseFistLower(statistics.ColumnName),
			FieldNameFL: domain.ToCamelCaseFistLarge(statistics.ColumnName),
			FieldType:   g.JavaConfigYaml.DbTypeMap[columnMap[statistics.ColumnName].DataType],
		})
	}

	for _, item := range indexMap {
		// sort index fields
		sort.Slice(item.Fields, func(i, j int) bool {
			return item.Fields[i].Index < item.Fields[j].Index
		})
		rs.IndexList = append(rs.IndexList, item)
	}
	return rs
}

func (g *JavaCodeGenerator) covertToModelStruct(tableInfo *mysql.TableInfo, params *CodeGenParams) *bo.ModelStruct {
	rs := &bo.ModelStruct{
		PackageName: params.PackageName,
		TableName:   tableInfo.Schema.TableName,
		BeanName:    domain.ToCamelCaseFistLarge(tableInfo.Schema.TableName),
		Columns:     make([]*bo.ModelField, 0),
		Comment:     tableInfo.Schema.TableComment,
	}

	// 填充fields
	for _, column := range tableInfo.Column {
		rs.Columns = append(rs.Columns, &bo.ModelField{
			ColumnName: column.ColumnName,
			Nullable:   column.IsNullable == "YES",
			IsPrimary:  column.ColumnKey == "PRI",

			FieldName: domain.ToCamelCaseFistLower(column.ColumnName),
			FieldType: g.JavaConfigYaml.DbTypeMap[column.DataType],
			Comment:   column.ColumnComment,
		})
	}
	return rs
}

type JavaConfigYaml struct {
	DbTypeMap map[string]string `yaml:"dbTypeMap"`
}
