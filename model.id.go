package main

import (
	"webapp/city"
	"webapp/code"
)

//National ID godoc
type NID struct {
	Nid     string `json:"nid"`
	Isvalid bool   `json:"isvalid"`
	City    string `json:"city"`
}

func getId(id string) (*NID, error) {
	isvalid := code.ValidateIranianCode(id)
	city := city.Detectcity(id)
	person := NID{Nid: id, Isvalid: isvalid, City: city}
	return &person, nil
}
