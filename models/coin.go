package models

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Data Model For Asset In Database
type CoinModel struct {
	ID              bson.ObjectId `json:"_id" bson:"_id"`
	Name            string        `json:"name" bson:"name"`
	Symbol          string        `json:"symbol" bson:"symbol"`
	Supply          uint64        `json:"supply" bson:"sypply"`
	ContractAddress string        `json:"contract_address" bson:"contract_address"`
	Created         time.Time     `json:"created" bson:"created"`
}

//var TemplateCollection *mgo.Collection

func getCollectionCoins() (error, *mgo.Collection) {
	err, db := GetDB()
	if err != nil {
		log.Printf("Error connecting db %v", err)
	}
	coinCollection := db.C("coins")
	return err, coinCollection
}

func CreateCoin(name string, symbol string, supply uint64, contract_address string) error {
	// TODO: save content to file
	// TODO: save model in db

	coin := CoinModel{
		bson.NewObjectId(),
		name,
		symbol,
		supply,
		contract_address,
		time.Time{},
	}
	_, coinCollection := getCollectionCoins()
	return coinCollection.Insert(coin)
}

func GetCoins(query interface{}) (error, []CoinModel) {
	var res []CoinModel
	_, coinCollection := getCollection()
	err := coinCollection.Find(query).All(&res)
	return err, res
}

//func GetById(id bson.ObjectId) (error, CoinModel) {
//	var res CoinModel
//	_, coinCollection := getCollection()
//	err := coinCollection.FindId(id).One(&res)
//	return err, res
//}
