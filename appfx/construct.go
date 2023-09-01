package appfx

import (
	"github.com/LiveAlone/go-util/domain/config"
	"github.com/LiveAlone/go-util/domain/template"
	"github.com/LiveAlone/go-util/service/api"
	"github.com/LiveAlone/go-util/service/api/yapi"
	"github.com/LiveAlone/go-util/service/model"
	"github.com/LiveAlone/go-util/service/model/lang"
)

func AppConstruct() []interface{} {
	depConstruct := []interface{}{
		config.NewConfigLoader, // 配置加载器
		UtilsLogger,            // 全局日志
	}

	// 支持命令行
	depConstruct = append(depConstruct, SubCmdConstructList()...)
	depConstruct = append(depConstruct, CommandProvider)

	// 模版生成器
	depConstruct = append(depConstruct, template.NewGenerator)

	// db 模型
	depConstruct = append(depConstruct, model.NewDaoGenerator, lang.NewCodeGenFactory, lang.NewJavaCodeGenerator, lang.NewGoCodeGenerator)

	// api gen
	depConstruct = append(depConstruct,
		api.NewSchemaGen,
		yapi.NewApiClient, // yapi api client
	)

	return depConstruct
}
