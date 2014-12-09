package routers

import (
	"github.com/astaxie/beego"

	"github.com/kempchee/washsales/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/uploads", &controllers.UploadsController{}, "get:Index")
	beego.Router("/upload", &controllers.UploadController{}, "post:CreateUpload")
	beego.Router("/download_csv/:csv_id", &controllers.UploadController{}, "*:DownloadCsv")
	beego.Router("/delete_csv/:csv_id", &controllers.UploadController{}, "*:DeleteCsv")
	beego.Router("/get_transactions/:csv_id", &controllers.UploadController{}, "*:ReadCsv")
}
