package MongoHelper

import (
	"log"
)

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Insert(DB string, Collection string, docs ...interface{}) {
	sessionDB, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	defer sessionDB.Close()
	c := sessionDB.DB(DB).C(Collection)
	err = c.Insert(&bson.M{"name": "Ale"})
	if err != nil {
		log.Println(err.Error())
		log.Fatal(err)
	}
}
