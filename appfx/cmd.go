package appfx

import (
	"github.com/LiveAlone/go-util/api/cmd"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// SubCmdList 注册添加命令
var SubCmdList = []interface{}{
	cmd.NewWordCmd,
	cmd.NewFileConvertCmd,
	cmd.NewCodeCmd,
	cmd.NewApiParam,
}

func SubCmdConstructList() (rs []any) {
	for _, v := range SubCmdList {
		rs = append(rs, fx.Annotate(
			v,
			fx.ResultTags(`group:"subCmd"`),
		))
	}
	return
}

type SubCmdListParam struct {
	fx.In
	SubCmdList []*cobra.Command `group:"subCmd"`
}

// CommandProvider 构建命令行
func CommandProvider(params SubCmdListParam) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "utils",
		Short: "utils",
		Long:  "个人项目工具",
	}
	for _, subCmd := range params.SubCmdList {
		rootCmd.AddCommand(subCmd)
	}
	return rootCmd, nil
}
