// Copyright 2015 Alexandre Cesar da Silva

package Domain

import (
	"time"
)

type Response struct {
	Date time.Time
	Data []string
}