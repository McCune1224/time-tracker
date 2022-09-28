package models

type User struct {
	Id        uint
	FirstName string
	LastName  string
	Username  string
	Email     string `gorm:"unique"`
	Password  string
}
