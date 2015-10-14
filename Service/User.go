// Copyright 2015 Alexandre Cesar da Silva

package Service

import (
	"alexcesar.tk/Gorms/DAO"
	"alexcesar.tk/Gorms/Domain"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func AddUser(user Domain.User) {
	user.Token = ""
	user.Forms = []Domain.Form {}
	user.Password = Hash(user.Password)

	collection := DAO.Collection("User")
	collection.Insert(user);
}

func RemoveUser(user Domain.User) {
	collection := DAO.Collection("User")
	collection.Remove(user)
}

func GetUserByToken(token string) Domain.User {
	collection := DAO.Collection("User")
	user := Domain.User{}
	collection.Find(bson.M {
		"token": token,
	}).One(&user)
	return user
}

func Login(email string, password string) string {
	password = Hash(password)
	collection := DAO.Collection("User")
	count, _ := collection.Find(bson.M {
		"email": email,
		"password": password,
	}).Count()
	if count != 1 {
		return ""
	}

	token := Hash(time.Now().String() + password)
	collection.Update(bson.M {
		"email": email,
		"password": password,
	}, bson.M { "$set": bson.M {
		"token": token,
	}})

	return token
}