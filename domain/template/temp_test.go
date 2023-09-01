package template

import (
	"bytes"
	"fmt"
	"github.com/LiveAlone/go-util/domain/template/bo"
	"testing"
	"text/template"
)

func TestNone(t *testing.T) {
	current, err := template.New("test").Funcs(nil).Parse("package {{ .PackageName }};")
	if err != nil {
		t.Error(err)
		return
	}

	var rs bytes.Buffer
	err = current.Execute(&rs, &bo.ModelStruct{
		PackageName: "com.test",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rs.String())
}
