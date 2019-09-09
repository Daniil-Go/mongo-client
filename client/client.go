package client

import "time"

//Client contains information about user
type Client struct {
	FirstName   string    `json:"firstName" bson:"firstName"`
	LastName    string    `json:"lastName" bson:"lastName"`
	Email       string    `json:"email" bson:"email"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	Login       string    `json:"login" bson:"login"`
	Password    string    `json:"password" bson:"password"`
	CreatedOn   time.Time `json:"createdOn" bson:"createdon"`
}
