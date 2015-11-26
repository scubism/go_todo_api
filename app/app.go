package app

import (
	"gopkg.in/mgo.v2"
	"os"
)

var DBSession *mgo.Session
var DB *mgo.Database

func InitApp() {
	initDB()
}

func CloseApp() {
	closeDB()
}

func initDB() {
	var err error
	DBSession, err = mgo.Dial(os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))
	if err != nil {
		panic(err)
	}

	DB = DBSession.DB(os.Getenv("MONGO_DB"))
}

func closeDB() {
	DBSession.Close()
}
