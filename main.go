package main

import (
	"fmt"
	"net/http"
)
func main() {
	fmt.Print("Go проект запущенный в Docker слушает на 9000 порту")
		handler := HttpHandler{}
		http.ListenAndServe(":9000", handler)
}

type HttpHandler struct{}
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello World!")
	res.Write(data)
}