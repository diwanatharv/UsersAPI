package models

type User struct {
	Id           int    `json:"Id" bson:"Id"` //this will uniquely identify each user
	FirstName    string `json:"FirstName"bson:"FirstName" validate:"required"`
	LastName     string ` json:"LastName" bson:"LastName" validate:"required"`
	Company      string `json:"Company" bson:"Company" validate:"required,min=1,max=20"`
	PhoneNumber  string `json:"PhoneNumber" bson:"PhoneNumber" validate:"required,len=10"`
	Email        string `json:"Email" bson:"Email" validate:"email,required"`
	Country      string ` json:"Country"bson:"Country" validate:"required"`
	BuisnessType string `json:"BuisnessType" bson:"BuisnessType"`
}
