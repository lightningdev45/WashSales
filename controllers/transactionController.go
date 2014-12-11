package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kempchee/washsales/connection"
	"github.com/kempchee/washsales/models"
	"gopkg.in/mgo.v2/bson"
)

type TransactionController struct {
	beego.Controller
}

func (this *TransactionController) UploadTransactions() {
	csvId := this.Ctx.Input.Param(":csv_id")
	var bsonCsvId bson.ObjectId
	if bson.IsObjectIdHex(csvId) {
		bsonCsvId = bson.ObjectIdHex(csvId)
	} else {
		bsonCsvId = bson.ObjectIdHex("")
	}
	transactions := []models.Transaction{}
	connection.TransactionsCollection.Find(bson.M{"uploadId": bsonCsvId}).All(&transactions)
	uploadsStruct := models.Transactions{transactions}
	this.Data["json"] = &uploadsStruct
	this.ServeJson()
}

func (this *TransactionController) Show() {
	transactionId := this.Ctx.Input.Param(":transactionId")
	var bsonTransactionId bson.ObjectId
	if bson.IsObjectIdHex(transactionId) {
		bsonTransactionId = bson.ObjectIdHex(transactionId)
	} else {
		bsonTransactionId = bson.ObjectIdHex("")
	}
	var transaction models.Transaction
	connection.TransactionsCollection.Find(bson.M{"_id": bsonTransactionId}).One(&transaction)
	this.Data["json"] = struct {
		Transaction *models.Transaction
	}{&transaction}
	this.ServeJson()
}
