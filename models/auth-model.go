package models

type DriverRegister struct {
	//Id        int    `json:"id"`
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

type UserRegister struct {
	//Id        int    `json:"id"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
	Token     string `json:"token"`
}

type SMS struct {
	Contact string `json:"sms"`
	Code    int    `json:"code"`
}
