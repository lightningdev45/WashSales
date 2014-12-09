package controllers

import (
	"encoding/csv"
	"github.com/astaxie/beego"
	"github.com/kempchee/washsales/connection"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Upload struct {
	File string        `json:"file"`
	Id   bson.ObjectId `bson:"_id" json:"id"`
}

type Uploads struct {
	Uploads []Upload `json:"uploads"`
}

type MainController struct {
	beego.Controller
}

type UploadsController struct {
	beego.Controller
}

type UploadController struct {
	beego.Controller
}

type Transaction struct {
	Ticker   string        `json:"ticker"`
	Action   string        `json:"action"`
	Quantity int64         `json:"quantity"`
	Price    float64       `json:"price"`
	Date     string        `json:"date"`
	Id       bson.ObjectId `bson:"_id" json:"id"`
	UploadId bson.ObjectId `bson:"uploadId" json:"upload"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

func (this *UploadsController) Index() {
	log.Println("hi")
	uploads := []Upload{}
	connection.UploadsCollection.Find(nil).All(&uploads)
	uploadsStruct := Uploads{uploads}
	log.Println(uploads)
	this.Data["json"] = &uploadsStruct
	this.ServeJson()

}

func (this *UploadController) CreateUpload() {
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
	newUpload := Upload{fileHeader.Filename, bson.NewObjectId()}
	connection.UploadsCollection.Insert(&newUpload)

	this.Data["json"] = newUpload
	this.ServeJson()
}

func (this *UploadController) ReadCsv() {
	csvId := this.Ctx.Input.Param(":csv_id")
	var bsonCsvId bson.ObjectId
	if bson.IsObjectIdHex(csvId) {
		bsonCsvId = bson.ObjectIdHex(csvId)
	} else {
		bsonCsvId = bson.ObjectIdHex("")
	}

	upload := Upload{}
	connection.UploadsCollection.Find(bson.M{"_id": bsonCsvId}).One(&upload)
	csvfile, err := os.Open("/home/kempchee/go/src/github.com/kempchee/washsales/private/" + upload.File)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for _, each := range rawCSVdata {
		log.Printf("email : %s and timestamp : %s\n", each[0], each[1])
	}

	var oneRecord Transaction

	var allRecords []Transaction

	for _, each := range rawCSVdata {
		oneRecord.Ticker = each[0]
		oneRecord.Date = each[1]
		oneRecord.Action = each[2]
		quantity, _ := strconv.ParseInt(each[3], 10, 64)
		price, _ := strconv.ParseFloat(each[4], 64)
		oneRecord.Quantity = quantity
		oneRecord.Price = price
		oneRecord.Id = bson.NewObjectId()
		oneRecord.UploadId = upload.Id
		allRecords = append(allRecords, oneRecord)
	}

	transactions := Transactions{allRecords}
	this.Data["json"] = &transactions
	this.ServeJson()
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
	csvId := this.Ctx.Input.Param(":csv_id")
	var bsonCsvId bson.ObjectId
	if bson.IsObjectIdHex(csvId) {
		bsonCsvId = bson.ObjectIdHex(csvId)
	} else {
		bsonCsvId = bson.ObjectIdHex("")
	}

	upload := Upload{}
	log.Println(csvId)
	connection.UploadsCollection.Find(bson.M{"_id": bsonCsvId}).One(&upload)
	log.Println(upload)
	log.Println(upload.File)
	this.Ctx.Output.ContentType("text/plain")
	this.Ctx.Output.Header("X-Accel-Redirect", "/private/"+upload.File)
	this.Ctx.Output.SetStatus(200)
}

func (c *MainController) Get() {
	c.TplNames = "index.tpl"
}
