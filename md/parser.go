package md

// Parser markdown解析器
type Parser interface {
	// Parse 通过解析器解析MD，
	Parse(md []byte) (html []byte)

	// MarkdownHTML 基于解析器生成的html字符串内容
	// 1. 读取目录文件，筛选包含以md文件结尾的markdown文件
	// 2. 并发的利用md解析lib来生成html源文件返回
	Markdown2HTML(mdfile string) (html string, err error)
}
