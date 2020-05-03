// go-md-render 基于输入的markdown文件，生产可执行的html，并启动一个简单的http服务器，渲染生产的html
package main

import (
	"flag"
	"fmt"
	"github.com/lupguo/mdrender/md"
	"github.com/lupguo/mdrender/render"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// Version 版本信息
const Version = "0.0.1"

var indexTpl, mdFile, cssFile, mdPath, listen string

// 接收参数：
// 	1. markdown所在目录
//	2. 利用blackfriday生成html文档在应用缓存中
//  3. 结合blackfriday生成html，与html模板组合，生成最终显示的html
//	4. 启动http服务器，指定ip端口，设定简单root文件http服务，提供http请求查看
func main() {
	flag.StringVar(&listen, "listen", "127.0.0.1:4500", "the listen address of the http server")
	flag.StringVar(&indexTpl, "tpl", "", "the template file for the html to be render, default is inner setting")
	flag.StringVar(&mdFile, "md", "README.md", "the markdown files to ben render")
	flag.StringVar(&mdPath, "path", "", "the static file path of the http server, default is the same as the path of the mdfile to be rendered")
	flag.StringVar(&cssFile, "css", "", "the css file using for render or using default css style")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Go Markdown Processor v"+Version+
			"\nUsage:\n"+
			"  %s [options] -md mdfile\n"+
			"Options:\n",
			os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	// 基于指定的目录，按html模板渲染内容
	http.HandleFunc("/index", Index)
	http.Handle("/", StaticHandler())

	log.Printf("listen on http://%s/index", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}

// StaticHandler 静态文件处理
func StaticHandler() http.Handler {
	var root string
	if mdPath == "" {
		root = path.Dir(mdFile)
	} else {
		root = mdPath
	}
	if _, err := os.Stat(root); err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.Dir(root))
}

// Index 首页处理器，基于markdown文件，结合静态文件目录，并渲染生成html响应
//	默认基于cmdline上配置的md，没有设置则为README.md
//	其次基于HTTP Request GET['mdfile']
func Index(w http.ResponseWriter, r *http.Request) {
	if f := r.FormValue("md"); f != "" {
		mdFile = f
	}
	html, err := index(indexTpl, mdFile)
	if err != nil {
		httpErrResponse(err, w)
		return
	}
	_, err = io.WriteString(w, html)
	if err != nil {
		httpErrResponse(err, w)
	}
}

func index(indexTpl, mdFile string) (html string, err error) {
	// body parser
	var p md.Parser
	p = md.NewBlackFriday()
	body, err := p.Markdown2HTML(mdFile)
	if err != nil {
		return "", err
	}

	// content builder
	css, err := render.MarkdownCSS(cssFile)
	if err != nil {
		return "", err
	}
	c := &render.Content{
		CSS:   css,
		Title: path.Base(mdFile),
		Body:  body,
	}

	// html render
	tmpl, err := render.NewTmpl(indexTpl)
	if err != nil {
		return "", err
	}
	return tmpl.Render(c)
}

func httpErrResponse(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}
