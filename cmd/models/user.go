package models

type User struct {
	BaseModel		 
	Email 	  string `gorm:"uniqueIndex;not null;size:255;"`
	Password  string `gorm:"not null;size:150;"`
}

func ToUserModel(email, pwd string) User {
	return User{Email: email, Password: pwd}
}