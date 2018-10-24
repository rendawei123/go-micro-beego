package main

import (
	"io"
	"net/http"
)

func main() {

	// 服务注册
	http.HandleFunc("/", Cookie)
	http.HandleFunc("/2", Cookie2)
	// 监听,第一个为地址，第二个为自己实现的handler
	http.ListenAndServe(":8080", nil)
}

// 传统设置方法
func Cookie(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello",
		Path:   "/",         // 设置路径
		Domain: "localhost", // 设置域名
		MaxAge: 120,         // 设置时间
	}

	// 将设置好的cookie出入
	http.SetCookie(w, ck)
	// 读取cookie
	ck2, err := r.Cookie("myCookie")

	// 如果有错误，打印错误内容
	if err != nil{
		io.WriteString(w, err.Error())
		return
	}

	// 如果没有错误，打印cookie的值  http: named cookie not present
	io.WriteString(w, ck2.Value)
}


func Cookie2(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello2",    // 注意value中不能有空格，为非法字符
		Path:   "/",         // 设置路径
		Domain: "localhost", // 设置域名
		MaxAge: 120,         // 设置时间
	}

	// 将设置好的cookie出入
	w.Header().Set("Set-Cookie", ck.String())
	// 读取cookie
	ck2, err := r.Cookie("myCookie")

	// 如果有错误，打印错误内容
	if err != nil{
		io.WriteString(w, err.Error())
		return
	}

	// 如果没有错误，打印cookie的值  http: named cookie not present
	io.WriteString(w, ck2.Value)
}
