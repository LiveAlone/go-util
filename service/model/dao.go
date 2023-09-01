package model

import (
	"fmt"
	"github.com/LiveAlone/go-util/domain/config"
	"github.com/LiveAlone/go-util/domain/mysql"
	"github.com/LiveAlone/go-util/service/model/lang"
	"log"
	"strings"
)

// DaoGenerator dao层代码生成器
type DaoGenerator struct {
	ConfigLoader *config.Loader
	Factory      *lang.CodeGenFactory
}

func NewDaoGenerator(configLoader *config.Loader, factory *lang.CodeGenFactory) *DaoGenerator {
	return &DaoGenerator{
		ConfigLoader: configLoader,
		Factory:      factory,
	}
}

// Gen 生成不同路径下代码
func (g *DaoGenerator) Gen(targetPath string) (rs map[string]string, err error) {
	// 1. 获取配置信息
	modelConfig := &Config{}
	err = g.ConfigLoader.LoadConfigToEntity(fmt.Sprintf("%s/%s", targetPath, "model.yaml"), modelConfig)
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
