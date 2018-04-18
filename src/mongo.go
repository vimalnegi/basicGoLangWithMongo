package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

var mongoURL = "localhost:27017"
var dataBaseName = "tada"

var mongoSession *mgo.Session

func getMongoDbFreshSession() (*mgo.Session, error) {
	return mgo.Dial(mongoURL)
}

// GetSession give the mongodb session
func GetSession() (*mgo.Session, error) {
	if mongoSession == nil {
		mongoSession, err := getMongoDbFreshSession()
		return mongoSession, err
	}
	clonedSession := mongoSession.Clone()
	return clonedSession, nil
}

// Create for creating generic
func Create(collectionName string, data interface{}) error {
	session, err := GetSession()
	if err != nil {
		return err
	}
	defer session.Close()
	collection := session.DB(dataBaseName).C(collectionName)
	return collection.Insert(data)
}

// FindOne From Collection
func FindOne(collectionName string, condition bson.M) (Demon, error) {
	session, err := GetSession()
	if err != nil {
		return Demon{}, err
	}
	defer session.Close()
	collection := session.DB(dataBaseName).C(collectionName)
	var result Demon
	fmt.Println(condition)
	err = collection.Find(condition).One(&result)
	fmt.Println("result")
	fmt.Println(result)
	return result, err
}
