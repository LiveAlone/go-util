package cmd

import (
	"fmt"
	"github.com/LiveAlone/go-util/service/model"
	"github.com/LiveAlone/go-util/util"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var targetPath string

func NewModelCmd(daoGenerator *model.DaoGenerator) *cobra.Command {
	desc := strings.Join([]string{
		"0. -d 指定生成目标地址",
		"1. 基于{target}/db.yaml获取代码生成配置信息",
		"2. 支持java/go 语言生成，默认java",
	}, "\n")

	cmd := &cobra.Command{
		Use:   "model",
		Short: "基于数据表生成Dao层代码",
		Long:  desc,
		Run: func(cmd *cobra.Command, args []string) {
			rs, err := daoGenerator.Gen(targetPath)
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
