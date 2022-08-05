package usermodel

import "go_restaurant/common"

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:roles;type:ENUM('user', 'admin')"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

//func (u *User) ComparePassword(hasher common.Hasher) bool {
//	hashedPassword := hasher.Hash()
//	return u.Password == hashedPassword
//}

func (u *User) IsActive() bool {
	if u == nil {
		return false
	}
	return u.Status == 1
}

//func (u *User) ToSimpleUser() *common.SimpleUser {
//	var simpleUser common.SimpleUser
//	simpleUser.ID = u.ID
//	simpleUser.Email = u.Email
//	simpleUser.Role = u.Role
//	return &simpleUser
//}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}
