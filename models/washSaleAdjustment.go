package models

import (
	"gopkg.in/mgo.v2/bson"
)

type WashSaleAdjustment struct {
	LotId          bson.ObjectId `bson:"lotId" json:"lot"`
	TransactionId  bson.ObjectId `bson:"transactionId" json:"transaction"`
	Quantity       int64         `bson:"quantity" json:"quantity"`
	TotalAmount    float64       `bson:"totalAmount" json:"totalAmount"`
	PerShareAmount float64       `bson:"perShareAmount" json:"perShareAmount"`
}

type WashSaleAdjustments struct {
	WashSaleAdjustments []WashSaleAdjustment
}
