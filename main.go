package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/kempchee/washsales/connection"
	_ "github.com/kempchee/washsales/routers"
	"log"
)

func main() {
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"

	log.Println("Starting Server")
	log.Println("Starting mongo db session")
	connection.Connect()
	defer connection.Session.Close()
	// Optional. Switch the session to a monotonic behavior.
	var AllowOtherOrigin = func(ctx *context.Context) {
		//origin := ctx.Input.Domain()
		ctx.Output.Header("Access-Control-Allow-Origin", "http://0.0.0.0:4200")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	}

	var ProperlyHandleOptionsRequest = func(ctx *context.Context) {
		if ctx.Input.Method() == "OPTIONS" {
			ctx.Output.Header("Status Code", "204")
			ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
			ctx.Output.Header("content-length", "0")
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, AllowOtherOrigin)
	beego.InsertFilter("/*", beego.BeforeRouter, ProperlyHandleOptionsRequest)
	beego.Run()
}
