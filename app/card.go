package main

//Card - Physical cards with QR codes
type Card struct {
	CardNo string `json:"cardno"`
	Active bool   `json:"active"`
}

//Cards are a thing
type Cards []Card
