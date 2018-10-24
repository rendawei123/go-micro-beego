package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hey</title>
</head>
<body>
<form method="post" action="/">
    Username: <input type="text" name="uname">
    Password: <input type="password" name="pwd">
    <button type="submit">Submit</button>
</form>
</body>
</html>
`

func main() {

	// 服务注册
	http.HandleFunc("/", Hey)
	// 监听,第一个为地址，第二个为自己实现的handler
	http.ListenAndServe(":8080", nil)
}

// 传统方法对表单进行解析
func Hey(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { // get方法让他返回模板
		// 创建模板对象
		t := template.New("hey")
		t.Parse(tpl)
		// 输出模版
		t.Execute(w, nil)
	} else {  // 否则视为post方法，即提交表单
		fmt.Println(r.FormValue("uname"), r.FormValue("pwd"))
	}
}
