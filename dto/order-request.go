package dto

type OrderRequest struct {
	UserId        int    `json:"userId"`
	LatitudeFrom  string `json:"latitudefrom"`
	LongitudeFrom string `json:"longitudeFrom"`
	LatitudeTo    string `json:"latitudeTo"`
	LongitudeTo   string `json:"longitudeTo"`
	Comments      string `josn:"comments"`
	Price         int    `json:"price"`
	Type          string `json:"type"`
	OrderStatus   int    `json:"orderStatus"`
}

type OrderResponse struct {
	Id            int    `json:"id"`
	UserId        int    `json:"userId"`
	LatitudeFrom  string `json:"latitudefrom"`
	LongitudeFrom string `json:"longitudeFrom"`
	LatitudeTo    string `json:"latitudeTo"`
	LongitudeTo   string `json:"longitudeTo"`
	Comments      string `josn:"comments"`
	Price         int    `json:"price"`
	Type          string `json:"type"`
	OrderStatus   int    `json:"orderStatus"`
}
