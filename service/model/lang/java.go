package lang

import (
	"fmt"
	"github.com/LiveAlone/go-util/domain"
	"github.com/LiveAlone/go-util/domain/config"
	"github.com/LiveAlone/go-util/domain/mysql"
	"github.com/LiveAlone/go-util/domain/template"
	"github.com/LiveAlone/go-util/domain/template/bo"
	"log"
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

	log.Println("modelContent:", modelContent)
	return
}

func (g *JavaCodeGenerator) convertIndexModeStruct(tableInfo *mysql.TableInfo, params *CodeGenParams) *bo.DaoStruct {
	rs := &bo.DaoStruct{
		PackageName: params.PackageName,
		TableName:   tableInfo.Schema.TableName,
		BeanName:    domain.ToCamelCaseFistLarge(tableInfo.Schema.TableName),
		Comment:     tableInfo.Schema.TableComment,
	}

	indexMap := make(map[string]*bo.DaoIndex)
	for _, statistics := range tableInfo.Index {
		// todo yqj fix 填充索引数据信息
		fmt.Println(statistics)
	}
	// 维护映射关系

	for _, item := range indexMap {
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
