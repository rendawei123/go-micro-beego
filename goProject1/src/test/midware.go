package main

import "net/http"

type SingleHost struct {
	handler     http.Handler
	allowedHost string
}

func main() {

}