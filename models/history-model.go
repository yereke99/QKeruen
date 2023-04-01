package models

type History struct {
	Id           int64 `json:"id"`
	OrderId      int   `json:"orderId"`
	DriverId     int   `json:"driverId"`
	UserId       int   `json:"userId"`
	StartDate    int   `json:"startDate"`
	FinishedDate int   `json:"finishedDate"`
}
