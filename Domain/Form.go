// Copyright 2015 Alexandre Cesar da Silva

package Domain

import (
	"time"
)

type Form struct {
	Name string
	Date time.Time
	Questions []Question
	Responses []Response
}