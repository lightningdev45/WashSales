package routers

import (
	"github.com/kempchee/chatrooms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/room/:roomId",&controllers.WebSocketController{})
}
