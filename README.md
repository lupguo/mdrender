## 精简的Markdown渲染器，支持Http Web Server 即时查看
> 类似于Hugo，快速指定一个markdown文件，利用内部markdown解析，然后通过HTTP服务请求查看被解析渲染的HTML文件
>
> 支持主模板`index.tpl`以及markdown样式文件指定，参见asset目录
>
> 可以利用HTTP GET请求更改需要被解析的markdown文件

## 使用方式

### 安装
```bash
go get -u github.com/lupguo/mdrender
```

### 快速执行
```
// 使用默认模板+样式渲染markdown文件，通过浏览器访问上述连接地址`http://127.0.0.1:4500/index`
$ mdrender -md ./README.md                                 
2020/05/03 19:27:34 listen on http://127.0.0.1:4500/index

// 更换一个主模板
$ mdrender -md ./README.md -tpl ./asset/index.tpl                                  
2020/05/03 19:40:18 listen on http://127.0.0.1:4500/index

// 更换模板、渲染样式文件
$ mdrender -md ../README.md -css ./asset/github-markdown.css -tpl ./asset/index.tpl

// 利用http get请求更换要被解析的markdown文件，在url中指定md在本地文件绝对路径
http://127.0.0.1:4500/index?md=/private/data/github.com/lupguo/go-example/gomemory/mdrender/README.md
```

### 命令行说明
```
$ mdrender -h         
Go Markdown Processor v0.0.1
Usage:
  mdrender [options] -md mdfile

Options:
  -css string
        the css file using for render or using default css style
  -listen string
        the listen address of the http server (default "127.0.0.1:4500")
  -md string
        the markdown files to ben render (default "README.md")
  -path string
        the static file path of the http server, default is the same as the path of the mdfile to be rendered
  -tpl string
        the template file for the html to be render, default is inner setting
```

### 截图
`$ mdrender -md ./README.md` 默认样式截图：

<img src="./asset/screen-shot.jpg" width="400"/>

`$ mdrender -md ../README.md -css ./asset/github-markdown.css -tpl ./asset/index.tpl` 指定模板、样式后的渲染效果：

<img src="./asset/screen-shot-2.jpg" width="400"/>

## 注意
当前应用程序，目前处理测试阶段，
