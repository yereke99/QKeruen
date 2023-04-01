package dto

type OfferRequest struct {
	//Id      int    `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Comment string `json:"comment"`
	Type    string `json:"type"`
}
