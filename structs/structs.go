package structs

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}

type Post struct{
	gorm.Model
	Title string
	Body string
	Author Person
}