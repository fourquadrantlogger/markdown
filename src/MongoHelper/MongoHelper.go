package MongoHelper

import (
	"log"
	"gopkg.in/mgo.v2"
	"encoding/json"
)

func Insert(DB string, Collection string,jsonstr string) {
	sessionDB, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	defer sessionDB.Close()
	c := sessionDB.DB(DB).C(Collection)
	var datamap map[string]interface{}
    if err := json.Unmarshal([]byte(jsonstr), &datamap); err == nil {
        log.Println("json str转map")
        log.Println(datamap)
    }	
	err = c.Insert(&datamap)
	if err != nil {
		log.Println(err.Error())
		log.Fatal(err)
	}
}
