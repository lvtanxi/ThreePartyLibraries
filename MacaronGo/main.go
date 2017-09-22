package main

import (
	"gopkg.in/macaron.v1"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ma := macaron.Classic()
	ma.Get("/",myHander)
	//用*号接受所有请求
	ma.Get("/*",myHander)
	log.Println("Server is runing")
	log.Println(http.ListenAndServe(":4040",ma))
}

func myHander(ctx *macaron.Context) string {
	return fmt.Sprintf("this request path is : %s", ctx.Req.RequestURI)
}
