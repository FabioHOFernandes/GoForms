// Copyright 2015 Alexandre Cesar da Silva

package DAO

import (
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func Collection(name string) *mgo.Collection {
	if db != nil {
		return db.C(name)
	}

	username := "Gorms"
	password := "qwe"
	host := "127.0.0.1"
	port := "27017"
	dbname := "Gorms"
	url := username + ":" + password + "@" + host + ":" + port + "/"
	url = host + ":" + port + "/"

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	db = session.DB(dbname)
	return Collection(name)
}