package models

import (
	"github.com/kempchee/washsales/connection"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	Id       bson.ObjectId `bson:"_id" json:"id"`
}

func GetUserByEmail(email string) User {
	user := User{}
	err := connection.UsersCollection.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		panic(err)
	}
	return user
}
