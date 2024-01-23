package core

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Mobile   int64
	Email    string
	Password string
}

// type Users struct {
// 	Users []User `json:"users"`
// }

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {

// 	// UUID VERSION 4
// 	user.ID = uuid.New()
// 	return

// }
