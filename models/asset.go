package models

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Data Model For Asset In Database
type AssetModel struct {
	ID               bson.ObjectId `json:"_id" bson:"_id"`
	Name             string        `json:"name" bson:"name"`
	Symbol           string        `json:"symbol" bson:"symbol"`
	Supply           uint64        `json:"supply" bson:"sypply"`
	Address          string        `json:"address" bson:"address"`
	Ownership        string        `json:"ownership" bson:"ownership"`
	OwnerSocialMedia string        `json:"owner_social_media" bson:"owner_social_media"`
	ProofOfTitle     string        `json:"proof_of_title" bson:"proof_of_title"`
	ContractAddress  string        `json:"contract_address" bson:"contract_address"`
	Created          time.Time     `json:"created" bson:"created"`
}

//var TemplateCollection *mgo.Collection

func getCollection() (error, *mgo.Collection) {
	err, db := GetDB()
	if err != nil {
		log.Printf("Error connecting db %v", err)
	}
	assetCollection := db.C("assets")
	return err, assetCollection
}

func CreateAsset(name string, symbol string, supply uint64, address string, ownership string, owner_social_media string, proof_of_title string, contract_address string) error {
	// TODO: save content to file
	// TODO: save model in db

	asset := AssetModel{
		bson.NewObjectId(),
		name,
		symbol,
		supply,
		address,
		ownership,
		owner_social_media,
		proof_of_title,
		contract_address,
		time.Time{},
	}
	_, assetCollection := getCollection()
	return assetCollection.Insert(asset)
}

func GetAssets(query interface{}) (error, []AssetModel) {
	var res []AssetModel
	_, assetCollection := getCollection()
	err := assetCollection.Find(query).All(&res)
	return err, res
}

func GetById(id bson.ObjectId) (error, AssetModel) {
	var res AssetModel
	_, assetCollection := getCollection()
	err := assetCollection.FindId(id).One(&res)
	return err, res
}
