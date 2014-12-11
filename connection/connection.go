package connection

import (
	"gopkg.in/mgo.v2"
)

var (
	Session                *mgo.Session
	UploadsCollection      *mgo.Collection
	TransactionsCollection *mgo.Collection
	err                    error
)

func Connect() error {
	Session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	UploadsCollection = Session.DB("washSales").C("uploads")
	TransactionsCollection = Session.DB("washSales").C("transactions")
	return err
}
