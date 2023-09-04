package code

import (
	"fmt"
	"github.com/LiveAlone/go-util/domain/config"
	"github.com/LiveAlone/go-util/domain/mysql"
	"github.com/LiveAlone/go-util/service/code/lang"
	"log"
	"strings"
)

// Generator dao层代码生成器
type Generator struct {
	ConfigLoader *config.Loader
	Factory      *lang.CodeGenFactory
}

func NewDaoGenerator(configLoader *config.Loader, factory *lang.CodeGenFactory) *Generator {
	return &Generator{
		ConfigLoader: configLoader,
		Factory:      factory,
	}
}

// GenClient 生成Rpc调用客户端代码
func (g *Generator) GenClient() (rs map[string]string, err error) {
	// 1. 获取配置信息
	clientConfig := &ClientConfig{}
	err = g.ConfigLoader.LoadConfigToEntity("client.yaml", clientConfig)
	if err != nil {
		return nil, err
	}

	// 2. 配置转换为代码schema

	// 3. 通过template 生成api 层接口协议

	return nil, err
}

// GenDao 生成dao持久层代码
func (g *Generator) GenDao(targetPath string) (rs map[string]string, err error) {
	// 1. 获取配置信息
	modelConfig := &Config{}
	err = g.ConfigLoader.LoadConfigToEntity(fmt.Sprintf("%s/%s", targetPath, "dao.yaml"), modelConfig)
	if err != nil {
		return nil, err
	}

	// 2. 获取sql模型
	analyser, err := mysql.NewTableSchemaAnalyser(modelConfig.Db.Url)
	if err != nil {
		return nil, err
	}

	tables := strings.Split(modelConfig.Db.Tables, ",")
	databaseName := modelConfig.Db.DataBase
	codeGen, err := g.Factory.GainGenerateFromLang(modelConfig.Target.Lang)
	if err != nil {
		return nil, err
	}

	rs = make(map[string]string)
	for _, table := range tables {
		tableInfo, err := analyser.TableInfo(databaseName, table)
		if err != nil {
			return nil, err
		}
		codes, err := codeGen.GenDaoFromTableInfo(tableInfo, &lang.CodeGenParams{
			PackageName: modelConfig.Target.Package,
		})
		if err != nil {
			log.Printf("generate dao error, info:%v err :%v", tableInfo, err)
			return nil, err
		}
		for k, v := range codes {
			rs[k] = v
		}
	}
	return
}
