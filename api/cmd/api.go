package cmd

import (
	"fmt"
	"log"

	"github.com/LiveAlone/go-util/domain"
	"github.com/LiveAlone/go-util/domain/config"
	"github.com/LiveAlone/go-util/domain/template"
	"github.com/LiveAlone/go-util/service/api"
	"github.com/LiveAlone/go-util/util"
	"github.com/spf13/cobra"
)

type ApiParam struct {
	project string
	allApi  bool
	list    string
	dest    string
}

var apiParam = new(ApiParam)

func NewApiParam(configLoader *config.Loader, gen *api.SchemaApiGen, templateGen *template.Generator) *cobra.Command {
	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "网关SDK生成",
		Run: func(cmd *cobra.Command, args []string) {
			// 初始化配置
			apiDestConfig := new(ApiConfig)
			err := configLoader.LoadConfigToEntity(fmt.Sprintf("%s/%s", apiParam.dest, "api.yaml"), apiDestConfig)
			if err != nil {
				log.Fatalf("yaml file read error %v", err)
			}
			generateFromApi(apiDestConfig.Token[apiParam.project], gen, templateGen)
		},
	}
	apiCmd.Flags().StringVarP(&apiParam.project, "project", "p", "", "输入需要生成项目")
	apiCmd.Flags().StringVarP(&apiParam.dest, "dest", "d", "", "输入目标文件路径")
	apiCmd.Flags().BoolVarP(&apiParam.allApi, "full", "f", false, "是否全量接口同步")
	apiCmd.Flags().StringVarP(&apiParam.list, "api", "a", "", "输入单个接口列表")
	return apiCmd
}

type ApiConfig struct {
	Token map[string]string `yaml:"token"`
}

func generateFromApi(token string, gen *api.SchemaApiGen, templateGen *template.Generator) {
	if len(token) == 0 {
		log.Fatalf("project fail get token, projet:%v", apiParam.project)
	}

	var content string
	var err error

	httpProject, err := gen.GenFromYapi(token, apiParam.allApi, apiParam.list)
	if err != nil {
		log.Fatalf("gen from yapi error, %v", err)
	}

	// dto generate
	dtoStructs := api.ConvertProjectApisDtoDesc(httpProject.ApiList)

	destBase := fmt.Sprintf("%s/api", apiParam.dest)
	_ = util.CreateAllParentDirs(destBase)

	//write dto
	content, err = templateGen.GenerateTemplateByName("ApiDto", map[string]any{
		"dtoList": dtoStructs,
	}, map[string]any{
		"ToCamelCaseFistLower": domain.ToCamelCaseFistLower,
	})
	if err != nil {
		log.Fatalf("generate template error, %v", err)
	}
	err = util.WriteFile(fmt.Sprintf("%s/%s_dto.go", destBase, httpProject.Name), []byte(content))
	if err != nil {
		log.Fatalf("wirte dto file error, %v", err)
	}

	// write client
	content, err = templateGen.GenerateTemplateByName("ApiClient", map[string]any{
		"apiList":  httpProject.ApiList,
		"basePath": httpProject.BasePath,
		"name":     domain.ToCamelCaseFistLarge(httpProject.Name),
	}, map[string]any{})
	if err != nil {
		log.Fatalf("generate template error, %v", err)
	}
	err = util.WriteFile(fmt.Sprintf("%s/%s_api.go", destBase, httpProject.Name), []byte(content))
	if err != nil {
		log.Fatalf("wirte client file error, %v", err)
	}

	// cont service
	for _, httpApi := range httpProject.ApiList {
		content, err = templateGen.GenerateTemplateByName("ApiControl", httpApi, map[string]any{})
		if err != nil {
			log.Fatalf("generate template error, %v", err)
		}
		err = util.WriteFile(fmt.Sprintf("%s/%s_%s_controller.go", destBase, domain.ToSnakeLower(httpApi.Prefix), httpProject.Name), []byte(content))
		if err != nil {
			log.Fatalf("wirte file error, %v", err)
		}

		content, err = templateGen.GenerateTemplateByName("ApiService", httpApi, map[string]any{})
		if err != nil {
			log.Fatalf("generate template error, %v", err)
		}
		err = util.WriteFile(fmt.Sprintf("%s/%s_%s_service.go", destBase, domain.ToSnakeLower(httpApi.Prefix), httpProject.Name), []byte(content))
		if err != nil {
			log.Fatalf("wirte file error, %v", err)
		}
	}
}
