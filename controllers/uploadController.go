package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kempchee/washsales/connection"
	"github.com/kempchee/washsales/models"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Index() {
	log.Println("hi")
	uploads := []models.Upload{}
	connection.UploadsCollection.Find(nil).All(&uploads)
	uploadsStruct := models.Uploads{uploads}
	log.Println(uploads)
	this.Data["json"] = &uploadsStruct
	this.ServeJson()
}

func (this *UploadController) CreateUpload() {
	if this.Ctx.Input.Method() == "OPTIONS" {
		this.ServeJson()
	} else {
		file, fileHeader, err := this.GetFile("file")
		if err != nil {
			panic(err)
		}
		fileBinary, newError := ioutil.ReadAll(file)
		if newError != nil {
			panic(err)
		}
		err = ioutil.WriteFile("/home/kempchee/go/src/github.com/kempchee/washsales/private/"+fileHeader.Filename, fileBinary, 0644)
		if err != nil {
			panic(err)
		}
		newUpload := models.Upload{fileHeader.Filename, bson.NewObjectId(), []bson.ObjectId{}}
		transactions := newUpload.ParseTransactionsFromUpload()
		var transactionIds []bson.ObjectId
		for _, v := range transactions {
			transactionIds = append(transactionIds, v.Id)
		}
		newUpload.TransactionIds = transactionIds
		connection.UploadsCollection.Insert(&newUpload)
		this.Data["json"] = newUpload
		this.ServeJson()
	}
}

func (this *UploadController) DeleteCsv() {
	csvId := this.Ctx.Input.Param(":csv_id")
	var bsonCsvId bson.ObjectId
	if bson.IsObjectIdHex(csvId) {
		bsonCsvId = bson.ObjectIdHex(csvId)
	} else {
		bsonCsvId = bson.ObjectIdHex("")
	}

	log.Println(csvId)
	err := connection.UploadsCollection.Remove(bson.M{"_id": bsonCsvId})
	if err != nil {
		panic(err)
	} else {
		this.Data["json"] = `{"success":"The File has been successfuly deleted."}`
		this.ServeJson()
	}
}

func (this *UploadController) DownloadCsv() {
	log.Println(this.Ctx)
	csvId := this.Ctx.Input.Param(":csv_id")
	var bsonCsvId bson.ObjectId
	if bson.IsObjectIdHex(csvId) {
		bsonCsvId = bson.ObjectIdHex(csvId)
	} else {
		bsonCsvId = bson.ObjectIdHex("")
	}

	upload := models.Upload{}
	log.Println(csvId)
	connection.UploadsCollection.Find(bson.M{"_id": bsonCsvId}).One(&upload)
	log.Println(upload)
	log.Println(upload.File)
	this.Ctx.Output.ContentType("text/plain")
	this.Ctx.Output.Header("X-Accel-Redirect", "/private/"+upload.File)
	this.Ctx.Output.SetStatus(200)
}
