package cmd

import (
	"fmt"
	"github.com/LiveAlone/go-util/service/code"
	"github.com/LiveAlone/go-util/util"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var targetPath string

func NewCodeCmd(generator *code.Generator) *cobra.Command {
	desc := strings.Join([]string{
		"0. -d 指定生成目标地址",
		"1. -t 指定代码生成类型，dao, client等",
		"1.1 type=dao 基于{target}/db.yaml获取代码生成配置信息",
		"2.1 type=client 基于{target}/client.yaml获取代码生成配置信息",
	}, "\n")

	cmd := &cobra.Command{
		Use:   "code",
		Short: "代码生成工具",
		Long:  desc,
		Run: func(cmd *cobra.Command, args []string) {
			rs, err := generator.GenDao(targetPath)
			if err != nil {
				log.Fatalf("generate dao error %v", err)
			}
			for path, content := range rs {
				targetPath := fmt.Sprintf("%s/%s", targetPath, path)
				err = util.CreateAllParentDirs(targetPath)
				if err != nil {
					log.Fatalf("create file parent dir fail :%v", err)
				}

				err = util.WriteFile(targetPath, []byte(content))
				if err != nil {
					log.Fatalf("wirte code file error :%v", err)
				}
			}
		},
	}
	cmd.Flags().StringVarP(&targetPath, "dest", "d", "", "文件生成目标地址")
	return cmd
}
