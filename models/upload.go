package models

import (
	"encoding/csv"
	"github.com/kempchee/washsales/connection"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"strconv"
)

type Upload struct {
	File           string          `bson:"file" json:"file"`
	Id             bson.ObjectId   `bson:"_id" json:"id"`
	TransactionIds []bson.ObjectId `bson:"transactionIds" json:"transactions"`
}

type Uploads struct {
	Uploads []Upload `json:"uploads"`
}

func (upload *Upload) ParseTransactionsFromUpload() []Transaction {

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
		log.Println("hi")

		log.Println(oneRecord)
		log.Println(&oneRecord)
		connection.TransactionsCollection.Insert(&oneRecord)
		allRecords = append(allRecords, oneRecord)
		log.Println("hi")
	}
	log.Println("hi")
	return allRecords
}
