package model

type User struct {
	Name          string  `json:"name" bson:"name"`
	Phoneno       int     `json:"no" bson:"no"`
	Profilestatus bool    `json:"status" bson:"ststus"`
	Balance       float64 `json:"balance" bson:"balance"`
}
type UserResponse struct {
	Data []User `json:"data"`
}
type ErrorResponse struct{
	Message string 
}
var Users = []User{
	{
		Name:          "GURU",
		Phoneno:       12345678,
		Profilestatus: true,
		Balance:       1345.899,
	},
	{
		Name:          "Jeeva",
		Phoneno:       1234560978,
		Profilestatus: true,
		Balance:       5545.899,
	},
}
