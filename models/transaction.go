package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	Ticker   string        `bson:"ticker" json:"ticker"`
	Action   string        `bson:"action" json:"action"`
	Quantity int64         `bson:"quantity" json:"quantity"`
	Price    float64       `bson:"price" json:"price"`
	Date     string        `bson:"date" json:"date"`
	Id       bson.ObjectId `bson:"_id" json:"id"`
	UploadId bson.ObjectId `bson:"uploadId" json:"upload"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}
