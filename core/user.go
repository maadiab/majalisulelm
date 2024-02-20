package core

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Mobile   int    `db:"mobile"`
	Email    string `db:"email"`
	Password string `db:"password"`
	// Permissions []pq.StringArray `db:"permissions"`
	Permissions string `db:"permissions"`
}

// type UserLogin struct
// type Users struct {
// 	Users []User `json:"users"`
// }

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {

// 	// UUID VERSION 4
// 	user.ID = uuid.New()
// 	return

// }
