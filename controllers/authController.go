package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/kempchee/washsales/connection"
	"github.com/kempchee/washsales/models"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
)

type AuthController struct {
	beego.Controller
}

var Store = sessions.NewCookieStore(
	[]byte(securecookie.GenerateRandomKey(64)), //Signing key
	[]byte(securecookie.GenerateRandomKey(32)))

var ErrInvalidPassword = errors.New("Invalid Password")

// Login attempts to login the user given a request.
func (this *AuthController) Login() {

	email, password := this.GetString("email"), this.GetString("password")
	session, _ := Store.Get(this.Ctx.Request, "washsalesSession")

	u := models.GetUserByEmail(email)

	//If we've made it here, we should have a valid user stored in u
	//Let's check the password
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Println("error")
		session.Values["id"] = nil
	} else {
		log.Println(u.Id.Hex())
		session.Values["id"] = u.Id.Hex()
		log.Println(session.Values["id"])
	}
	session.Save(this.Ctx.Request, this.Ctx.ResponseWriter)
	u.Password = ""
	this.Data["json"] = struct {
		User models.User `json:"user"`
	}{u}
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	this.ServeJson()
}

// Register attempts to register the user given a request.
func (this *AuthController) Register() {

	email, password := this.GetString("email"), this.GetString("password")
	//u, err := models.GetUserByemail(email)
	// If we have an error which is not simply indicating that no user was found, report it
	//if err != sql.ErrNoRows {
	//return false, err
	//}
	//If we've made it here, we should have a valid email given
	//Let's create the password hash
	log.Println(email, password)
	var u models.User
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Email = email
	u.Password = string(h)
	u.Id = bson.NewObjectId()
	if err != nil {
		panic(err)
	}
	err = connection.UsersCollection.Insert(&u)
	if err != nil {
		panic(err)
	}
	u.Password = ""
	this.Data["json"] = struct {
		User models.User `json:"user"`
	}{u}
	this.ServeJson()
}

func (this *AuthController) CurrentUser() {
	session, err := Store.Get(this.Ctx.Request, "washsalesSession")
	if err != nil {
		//panic(err)
	}
	UserId := session.Values["id"]
	if id, ok := UserId.(string); ok {
		var bsonUserId bson.ObjectId
		if bson.IsObjectIdHex(id) {
			bsonUserId = bson.ObjectIdHex(id)
		} else {
			bsonUserId = bson.ObjectIdHex("")
		}
		var user models.User
		err := connection.UsersCollection.Find(bson.M{"_id": bsonUserId}).One(&user)
		if err != nil {
			panic(err)
		}
		user.Password = ""
		this.Data["json"] = struct {
			User models.User `json:"user"`
		}{user}
	}

	this.ServeJson()
}

func (this *AuthController) Logout() {
	session, err := Store.Get(this.Ctx.Request, "washsalesSession")
	if err != nil {
		//panic(err)
	}
	session.Values["id"] = nil
	session.Save(this.Ctx.Request, this.Ctx.ResponseWriter)
	this.Data["json"] = struct {
	}{}
	this.ServeJson()
}

func GenerateSecureKey() string {
	// Inspired from gorilla/securecookie
	k := make([]byte, 32)
	io.ReadFull(rand.Reader, k)
	return fmt.Sprintf("%x", k)
}
