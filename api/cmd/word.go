package cmd

import (
	"github.com/LiveAlone/go-util/domain"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper                = iota + 1 // 全部转大写
	ModeLower                           // 全部转小写
	ModeToCamelCaseFistLarge            // 转大写驼峰
	ModeToCamelCaseFistLower            // 转小写驼峰
	ModeToSnakeLower                    // 转小写下划线
	ModeToSnakeLarge                    // 转大写下划线
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：转大写驼峰",
	"4：转小写驼峰",
	"5：转下划线小写",
	"6：转下划线大写",
}, "\n")

type WordParam struct {
	str  string
	mode int8
}

var wordParam = new(WordParam)

func NewWordCmd() *cobra.Command {

	wordCmd := &cobra.Command{
		Use:   "word",
		Short: "单词格式转换",
		Long:  desc,
		Run: func(cmd *cobra.Command, args []string) {
			var content string
			var str = wordParam.str
			switch wordParam.mode {
			case ModeUpper:
				content = domain.ToUpper(str)
			case ModeLower:
				content = domain.ToLower(str)
			case ModeToCamelCaseFistLarge:
				content = domain.ToCamelCaseFistLarge(str)
			case ModeToCamelCaseFistLower:
				content = domain.ToCamelCaseFistLower(str)
			case ModeToSnakeLower:
				content = domain.ToSnakeLower(str)
			case ModeToSnakeLarge:
				content = domain.ToSnakeLarge(str)
			default:
				log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
			}

			log.Printf("输出结果: %s", content)
		},
	}
	wordCmd.Flags().StringVarP(&wordParam.str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&wordParam.mode, "mode", "m", 0, "请输入单词转换的模式")
	return wordCmd
}
