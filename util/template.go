package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

// GenerateFromTemplate 模版生成文本内容
func GenerateFromTemplate(templateName string, data any, funcMap template.FuncMap) string {
	filePath := fmt.Sprintf("conf/template/%s.template", templateName)
	bc, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("os read file error, %s, %v", filePath, err)
	}

	templateContent := string(bc)
	current, err := template.New("current").Funcs(funcMap).Parse(templateContent)
	if err != nil {
		log.Fatalf("template compile error, content:%s, cause:%v ", templateContent, err)
	}

	var rs bytes.Buffer
	err = current.Execute(&rs, data)
	if err != nil {
		log.Fatalf("template execute error, data:%v, tmp:%v, cause:%v", data, templateName, err)
	}
	return rs.String()
}
