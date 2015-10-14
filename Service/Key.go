// Copyright 2015 Alexandre Cesar da Silva

package Service

import (
	"alexcesar.tk/Gorms/DAO"
	"crypto/sha1"
	"encoding/base64"
	"gopkg.in/mgo.v2/bson"
)

func salt() string {
	return "500753244"
}

func Hash(value string) string {
	hash := sha1.New()
	bytes := hash.Sum([]byte(value + salt()))
	return base64.URLEncoding.EncodeToString(bytes)
}

func GetPublicKey() string {
	collection := DAO.Collection("Keys")

	public := bson.M {}
	collection.Find(bson.M {
		"name": "public",
	}).One(&public)

	result, ok := public["value"].(string)
	if ok {
		return result
	} else {
		return ""
	}
}

func GetPrivateKey() string {
	collection := DAO.Collection("Keys")

	private := bson.M {}
	collection.Find(bson.M {
		"name": "private",
	}).One(&private)

	result, ok := private["value"].(string)
	if ok {
		return result
	} else {
		return ""
	}
}