package lang

import (
	"errors"
	"fmt"
	"github.com/LiveAlone/go-util/domain/mysql"
)

// CodeGen 不同代码语言生成器
type CodeGen interface {
	// GenDaoFromTableInfo 通过数据表信息生成dao层代码
	GenDaoFromTableInfo(tableInfo *mysql.TableInfo, params *CodeGenParams) (rs map[string]string, err error)
}

// CodeGenParams 代码生成相关额外参数
type CodeGenParams struct {
	PackageName string // com.xxx.xxx
}

// CodeGenFactory 代码生成器工厂
type CodeGenFactory struct {
	Java *JavaCodeGenerator
	Go   *GoCodeGenerator
}

func NewCodeGenFactory(java *JavaCodeGenerator, goLang *GoCodeGenerator) *CodeGenFactory {
	return &CodeGenFactory{
		Java: java,
		Go:   goLang,
	}
}

func (c *CodeGenFactory) GainGenerateFromLang(lang string) (CodeGen, error) {
	switch lang {
	case "java":
		return c.Java, nil
	case "go":
		return c.Go, nil
	default:
		return nil, errors.New(fmt.Sprintf("not support lang: %s", lang))
	}
}
