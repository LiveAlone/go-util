package cmd

import (
	"log"

	"github.com/LiveAlone/go-util/util"
	"github.com/spf13/cobra"
)

type FileConvertParam struct {
	fromFile string
	destFile string
}

var fileConvertParam = new(FileConvertParam)

// 1. list 选择任务taskId 列表
// 2. source 定义依赖来源， 任务分类管理。 工作注解方式，注入判断。
// 3. func 数据组装合并展示结果。
// 4. fmt, excel, file等等。

// NewFileConvertCmd 文件行转换
func NewFileConvertCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert",
		Short: "文件行转换",
		Long:  "用户需求进行文件行格式化",
		Run: func(cmd *cobra.Command, args []string) {
			lines, err := util.ReadFileLines(fileConvertParam.fromFile)
			if err != nil {
				log.Fatal(err)
				return
			}

			destLines := make([]string, 0, len(lines))
			for i, line := range lines {
				destLine, exit := lineConvert(i, line)
				destLines = append(destLines, destLine)
				if exit {
					break
				}
			}
			err = util.WriteFileLines(fileConvertParam.destFile, destLines)
			if err != nil {
				log.Fatal(err)
				return
			}
		},
	}

	cmd.Flags().StringVarP(&fileConvertParam.fromFile, "from", "f", "from.text", "文件来源")
	cmd.Flags().StringVarP(&fileConvertParam.destFile, "to", "t", "to.text", "文件输出地址")

	return cmd
}

func lineConvert(i int, line string) (string, bool) {
	return line + "ccc", false
}
