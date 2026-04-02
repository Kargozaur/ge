package models

type User struct {
	BaseModel		 
	Email 	  string `gorm:"uniqueIndex;not null;size:255;"`
	Password  string `gorm:"not null;size:150;"`
}