package models

type DriverModel struct {
	Id        int64  `json:"id"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
	CarNumber string `json:"carNumber"`
	CarColor  string `json:"carColor"`
	CarModel  string `json:"carModel"`
	DocsFront string `json:"docsFront"`
	DocsBacks string `json:"docsBacks"`
	CarType   string `json:"carType"`
	Token     string `json:"token"`
}
