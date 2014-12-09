package main

import (
	"github.com/astaxie/beego"
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

	beego.Run()
}
