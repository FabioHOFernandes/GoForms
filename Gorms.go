package main

import (
//	"huijari.ml/Gorms/Domain"
	"huijari.ml/Gorms/Service"
)

func main() {
	user := Service.GetUserByToken("MjAxNS0xMC0xMyAyMTowMjo1NC4xOTk2NDExMzYgLTAzMDAgQlJUY1hkbE1USXpOVEF3TnpVek1qUTAyam1qN2w1clN3MHlWYl92bFdBWWtLX1lCd2s9NTAwNzUzMjQ02jmj7l5rSw0yVb_vlWAYkK_YBwk=")
	forms := Service.GetUserForms(user)

	for _, form := range forms {
		name, ok := form["name"].(string)
		if ok {
			print(name + "\n")
		}
	}
}