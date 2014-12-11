package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Lot struct {
	Ticker        string        `bson:"ticker" json:"ticker"`
	Quantity      int64         `bson:"quantity" json:"quantity"`
	Price         float64       `bson:"price" json:"price"`
	DateAcquired  string        `bson:"date" json:"date"`
	Id            bson.ObjectId `bson:"_id" json:"id"`
	TransactionId bson.ObjectId `bson:"transactionId" json:"transaction"`
}

type Lots struct {
	Lots []Lot `json:"lots"`
}
