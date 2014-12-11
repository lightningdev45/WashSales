package models

import (
	"gopkg.in/mgo.v2/bson"
)

type LotSale struct {
	LotId         Lot           `bson:"lotId" json:"lot"`
	Id            bson.ObjectId `bson:"_id" json:"id"`
	SaleDate      string        `bson:"date" json:"date"`
	TransactionId bson.ObjectId `bson:"transactionId" json:"transaction"`
	Quantity      int64         `bson:"quantity" json:"quantity"`
	Price         float64       `bson:"price" json:"price"`
}

type LotSales struct {
	LotSales []LotSale `json:"uploads"`
}
