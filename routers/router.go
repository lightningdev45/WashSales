package routers

import (
	"github.com/astaxie/beego"

	"github.com/kempchee/washsales/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/uploads", &controllers.UploadController{}, "get:Index")
	beego.Router("/upload", &controllers.UploadController{}, "*:CreateUpload")
	beego.Router("/download_csv/:csv_id", &controllers.UploadController{}, "*:DownloadCsv")
	beego.Router("/delete_csv/:csv_id", &controllers.UploadController{}, "*:DeleteCsv")
	beego.Router("/get_transactions/:csv_id", &controllers.TransactionController{}, "*:UploadTransactions")
	beego.Router("/transactions/:transactionId", &controllers.TransactionController{}, "get:Show")
	beego.Router("/users_create", &controllers.AuthController{}, "post:Register")
	beego.Router("/sign_in", &controllers.AuthController{}, "post:Login")
	beego.Router("/current_user", &controllers.AuthController{}, "get:CurrentUser")
	beego.Router("/sign_out", &controllers.AuthController{}, "post:Logout")
}
