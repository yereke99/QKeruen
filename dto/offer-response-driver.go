package dto

type OfferResponseDriver struct {
	Id      int64  `json:"id"`
	To      string `json:"to"`
	From    string `json:"from"`
	Comment string `json:"comment"`
	Driver  int    `json:"driver"` // there is mistake, so you must to change!
}

type OfferResponseUser struct {
	Id      int64  `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Comment string `json:"comment"`
	Type    string `json:"type"`
	User    int    `json:"user"` // there is mistake, so you must to change!
}
