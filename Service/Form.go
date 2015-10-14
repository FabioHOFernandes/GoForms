// Copyright 2015 Alexandre Cesar da Silva

package Service

import (
	"alexcesar.tk/Gorms/DAO"
	"alexcesar.tk/Gorms/Domain"
	"gopkg.in/mgo.v2/bson"
)

func AddForm(user Domain.User, form Domain.Form) {
	collection := DAO.Collection("User")
	collection.Update(user, bson.M { "$set": bson.M {
		"forms": append(user.Forms, form),
	}});
}

func GetUserForms(user Domain.User) []bson.M {
	forms := []bson.M {}
	for _, form := range user.Forms {
		forms = append(forms, bson.M {
			"name": form.Name,
			"date": form.Date,
		})
	}
	return forms
}

func GetForm(user Domain.User, index int) bson.M {
	form := user.Forms[index]
	return bson.M {
		"name": form.Name,
		"date": form.Date,
		"questions": form.Questions,
	}
}

func GetFormResponses(user Domain.User, index int) []Domain.Response {
	form := user.Forms[index]
	return form.Responses
}

func AddResponse(user Domain.User, index int, response Domain.Response) {
	// TODO!!!
}