package models

type OfferUserModel struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Comment string `json:"comment"`
	Type    string `json:"type"`
	Driver  int    `json:"driver"`
}
