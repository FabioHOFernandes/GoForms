// Copyright 2015 Alexandre Cesar da Silva

package Domain

type User struct {
	Email string
	Password string
	Token string
	Forms []Form
}