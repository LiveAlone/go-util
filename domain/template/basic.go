// 通过机构体类型维护映射关系

package template

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"text/template"
)

var structNameTemplateMap = map[string]string{
	//"HelloTemplate": "hello",
	////"ModelStruct":   "model/flow",
	//"ModelStruct": "model/basic",
	//"DataStruct":  "model/data",
	//
	//"ApiDto":     "api/dto",
	//"ApiClient":  "api/client",
	//"ApiControl": "api/control",
	//"ApiService": "api/service",

	"JpaModel": "java/dao/model_jpa",
}

// Generator 模版生成器
type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateTemplateByName(templateName string, data any, funcMap template.FuncMap) (string, error) {
	templatePath, ok := structNameTemplateMap[templateName]
	if !ok {
		return "", fmt.Errorf("template not found, struct:%s", templateName)
	}

	filePath := fmt.Sprintf("conf/template/%s.template", templatePath)
	bc, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	templateContent := string(bc)
	current, err := template.New(templateName).Funcs(funcMap).Parse(templateContent)
	if err != nil {
		return "", err
	}

	var rs bytes.Buffer
	err = current.Execute(&rs, data)
	if err != nil {
		return "", err
	}
	return rs.String(), nil
}

func (g *Generator) GenerateTemplateContent(data any, funcMap template.FuncMap) (string, error) {
	var dataStructName string
	dataType := reflect.TypeOf(data)
	switch dataType.Kind() {
	case reflect.Ptr:
		dataStructName = dataType.Elem().Name()
	case reflect.Struct:
		dataStructName = dataType.Name()
	default:
		return "", fmt.Errorf("data type not support, type:%v", dataType.Kind())
	}

	return g.GenerateTemplateByName(dataStructName, data, funcMap)
}
