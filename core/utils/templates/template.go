package templates

import (
	"bytes"
	"github.com/treeyh/soc-go-common/core/errors"
	"text/template"
)

// Render 模板解析工具
func Render(str string, params interface{}) (string, errors.AppError) {
	tmpl, err := template.New("test").Parse(str) //建立一个模板
	if err != nil {
		return "", errors.NewAppErrorByExistError(errors.TemplateRenderFail, err)
	}

	buf := bytes.NewBufferString("")

	err = tmpl.Execute(buf, params) //将struct与模板合成，合成结果放到os.Stdout里
	if err != nil {
		return "", errors.NewAppErrorByExistError(errors.TemplateRenderFail, err)
	}
	return buf.String(), nil
}
