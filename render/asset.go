package render


var indexTpl = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.Title}}</title>
		<style type="text/css">
		{{.CSS}}
		</style>
</head>
<body>
<article class="markdown-body">
{{.Body}}
</article>
</body>
</html>
`

var defaultStyle = `
body{margin:0 auto;font-family:ubuntu,Tahoma,"Microsoft YaHei",arial,sans-serif;color:#444;line-height:1;padding:30px}img{max-width:100%}@media screen and (min-width:1000px){body{width:1024px;margin:10px auto}}h1,h2,h3,h4{color:#111;font-weight:400;margin-top:1em}h1,h2,h3,h4,h5{font-family:Georgia,Palatino,serif}dl,h1,h2,h3,h4,h5{margin-bottom:16px;padding:0}p{margin-top:8px;margin-bottom:3px}h1{font-size:48px;line-height:54px}h2{font-size:36px;line-height:42px}h1,h2{border-bottom:1px solid #efeaea;padding-bottom:10px}h3{font-size:24px;line-height:30px}h4{font-size:21px;line-height:26px}h5{font-size:18px;line-height:23px}a{color:#09f;margin:0 2px;padding:0;vertical-align:baseline;text-decoration:none}a:hover{text-decoration:none;color:#f60}ol,ul{padding:0;padding-left:18px;margin:0}li{line-height:24px}ol,p,ul{font-size:16px;line-height:24px}ol ol,ul ol{list-style-type:lower-roman}code,pre{font-family:Consolas,Monaco,Andale Mono,monospace;background-color:#f7f7f7;color:inherit}code{font-family:Consolas,Monaco,Andale Mono,monospace;margin:0 2px}pre{font-family:Consolas,Monaco,Andale Mono,monospace;line-height:1.7em;overflow:auto;padding:6px 10px;border-left:5px solid #6ce26c}pre>code{font-family:Consolas,Monaco,Andale Mono,monospace;border:0;display:inline;max-width:initial;padding:0;margin:0;overflow:initial;line-height:1.6em;font-size:.95em;white-space:pre;background:0 0}code{color:#666555}aside{display:block;float:right;width:390px}blockquote{border-left:.5em solid #eee;padding:0 0 0 2em;margin-left:0}blockquote cite{font-size:14px;line-height:20px;color:#bfbfbf}blockquote cite:before{content:'\2014 \00A0'}blockquote p{color:#666}hr{text-align:left;color:#999;height:2px;padding:0;margin:16px 0;background-color:#e7e7e7;border:0 none}dl{padding:0}dl dt{padding:10px 0;margin-top:16px;font-size:1em;font-style:italic;font-weight:700}dl dd{padding:0 16px;margin-bottom:16px}dd{margin-left:0}table{border-spacing:0;width:100%}table{border:solid #ccc 1px}table thead{background:#f7f7f7}table thead tr:hover{background:#f7f7f7}table tr:hover{background:#fbf8e9;-o-transition:all .1s ease-in-out;-webkit-transition:all .1s ease-in-out;-moz-transition:all .1s ease-in-out;-ms-transition:all .1s ease-in-out;transition:all .1s ease-in-out}.table th,table td{border-left:1px solid #ccc;border-top:1px solid #ccc;padding:10px;text-align:left}table th{border-top:none;text-shadow:0 1px 0 rgba(255,255,255,.5);padding:5px;border-left:1px solid #ccc}table td:first-child,table th:first-child{border-left:none}
`
