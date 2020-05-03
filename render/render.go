package render

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

// NewTmpl 初始化Markdown渲染模板
func NewTmpl(f string) (*tmpl, error) {
	if f == "" {
		return defaultTmpl, nil
	}
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return &tmpl{string(file)}, nil
}


// tmpl 渲染模板
type tmpl struct {
	content string
}

var defaultTmpl = &tmpl{indexTpl}

// HTML 待渲染的HTML实例
type Content struct {
	CSS   string
	Title string
	Body  string
}

// Render 基于模板结合HTML实例数据，渲染HTML
func (tpl *tmpl) Render(c *Content) (html string, err error) {
	t := template.Must(template.New("md").Parse(tpl.content))

	var buf = &bytes.Buffer{}
	err = t.Execute(buf, c)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// MarkdownCSS 返回渲染markdown用的默认css样式
func MarkdownCSS(f string) (style string, err error) {
	if f == "" {
		return defaultStyle, nil
	}
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
