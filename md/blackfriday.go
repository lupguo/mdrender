package md

import (
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"strings"
)

// BlackFriday Markdown解析器 https://github.com/russross/blackfriday
type BlackFriday struct {
	Option blackfriday.Extensions
}

// NewBlackFriday 初始化一个BlackFriday解析器
func NewBlackFriday() *BlackFriday {
	return &BlackFriday{}
}

// Parse 通过BlackFriday解析器解析MD
func (p *BlackFriday) Parse(input []byte) (html []byte) {
	return blackfriday.Run(input)
}

// Markdown2HTML
func (p *BlackFriday) Markdown2HTML(mdfile string) (html string, err error) {
	// 读取md文件
	b, err := ioutil.ReadFile(mdfile)
	if err != nil {
		return "", err
	}

	// 解析md文件
	var s = strings.Builder{}
	s.Write(p.Parse(b))
	return s.String(), nil
}
