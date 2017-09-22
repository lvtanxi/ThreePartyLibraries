package main

import (
	"gopkg.in/macaron.v1"
	"fmt"
	"log"
	"os"
)

//创建全局服务

var logger = log.New(os.Stdout, "[App]", 0)

func main() {
	ma := macaron.Classic()

	//创建全局服务
	ma.Map(logger)

	/*	//在用一个请求中，多个处理器之间相互传递参数
		ma.Get("/", func(ctx *macaron.Context) {
			ctx.Data["Num"]=1
		}, func(ctx *macaron.Context) {
			ctx.Data["Num"]=ctx.Data["Num"].(int)+1
		}, func(ctx *macaron.Context) string{
			return fmt.Sprintf("num : %d",ctx.Data["Num"])
		})*/
	//利用Query获取url地址中的参数,
	/*	ma.Get("/", func(ctx *macaron.Context)string {
			return fmt.Sprintf("Uid 64 :%d",ctx.QueryInt64("uid"))
		})*/
	//获取远程IP
	/*
	ma.Get("/",  func(ctx *macaron.Context)string{
		return ctx.RemoteAddr()
	})
	*/
	//Next让出处理器,让其他Hander先执行
	ma.Get("/next", next1, next2, next3)
	//设置全局cookie加密
	/*	ma.SetDefaultCookieSecret("macaron")
		ma.Get("/setCo", func(ctx *macaron.Context)string {
			//普通cookie
			ctx.SetCookie("user","jaychou")
			//加密cookie
			ctx.SetSecureCookie("lv","jaychou")
			//超级加密
			ctx.SetSuperSecureCookie("lv","wuwuwu","jaychou")
			return "设置成功了"
		})

		ma.Get("/getCo", func(ctx *macaron.Context)string {
			val,_:=ctx.GetSecureCookie("lv")
			val2,_:=ctx.GetSuperSecureCookie("lv","wuwuwu")
			return fmt.Sprintf("%s : %s : %s",ctx.GetCookie("user"),val,val2)
		})*/
	//ctx.Resp.Write只有返回了，请求就不再继续
	/*
	ma.Get("/xxx", func(ctx *macaron.Context) {
			ctx.Data["Count"] = 1
		},
		func(ctx *macaron.Context) {
			ctx.Data["Count"] = ctx.Data["Count"].(int) + 1
		},
		func(ctx *macaron.Context) {
			ctx.Data["Count"] = ctx.Data["Count"].(int) + 1
		},
		func(ctx *macaron.Context) {
			ctx.Resp.Write([]byte("你好！实际"))
		},
		func(ctx *macaron.Context) string{
			return fmt.Sprintf("There are %d handlers before this", ctx.Data["Count"])
		},
		func(ctx *macaron.Context) string {
			return fmt.Sprintf("this is last,There are %d handlers before this", ctx.Data["Count"])
		},
	)
	*/
	//容错恢复，在正式的时候会隐藏
	/*	ma.Get("/panic", func(l *log.Logger) {
			//全局日志
			l.Println("这是一行日子")

			panic("有钱就是认识")
		})*/

	/*
	ma.Get("/", func(l *log.Logger) string {
		l.Println("这是我的全局日志器")
		return "全局服务"
	})
	ma.Get("/test", func(l *log.Logger) string {
		l.Println("这是我的全局日志器")
		return "全局服务"
	})
	*/

	ma.Get("/", func(l *log.Logger) string {
		l.Println("这是我的全局日志器")
		return "全局服务"
	})
	ma.Get("/test",myLogger, func(l *log.Logger) string {
		l.Println("这是我的请求级别日志器")
		return "这是我的请求级别日志器"
	})
	/*
	ma.Get("/testx",myLoaclLogger, func(l log.Logger) string {
		l.Println("映射值到接口")
		return "映射值到接口"
	})
	*/


	ma.Run()

}

func myLoaclLogger(ctx *macaron.Context)  {
	logger := log.New(os.Stdout, "[xxxx]", 0)
	ctx.MapTo(logger,(Logger)(nil))
}

type Logger interface {
	Println(... interface{})
}


func next1(ctx *macaron.Context) {
	fmt.Println("处理器1进入")
	ctx.Next()
	fmt.Println("处理器1退出")
}
func next2(ctx *macaron.Context) {
	fmt.Println("处理器2进入")
	ctx.Next()
	fmt.Println("处理器2退出")
}
func next3(ctx *macaron.Context) {
	fmt.Println("处理器3进入")
	ctx.Next()
	fmt.Println("处理器3退出")
}

func myLogger(ctx *macaron.Context){
	logger := log.New(os.Stdout, "[hhhh]", 0)
	ctx.Map(logger)
}
