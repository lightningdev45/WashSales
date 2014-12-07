package main

import (
	_ "github.com/kempchee/chatrooms/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"

	beego.Run()
}
