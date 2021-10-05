package db

import (
	"finbuild/entity/finance"
	"finbuild/entity/user"
	"github.com/google/uuid"
)

func UserRegistry(user *user.User, wallet *finance.Wallet) {

	DB.Model(&user).Create(&user)

	DB.Table("wallets").Model(&finance.Wallet{}).Create(&wallet)

}

// GetUser get all information from users where userID
func GetUser(u uuid.UUID) *user.User {

	uu := user.User{}
	DB.First(&uu, u).Scan(&uu)
	return &uu
}

// GetShortUser get short information from users where userID
func GetShortUser(u uuid.UUID) *user.ShortUsers {

	su := user.ShortUsers{}
	DB.Table("users").First(&su, "user_id = ?", u).Scan(&su).Find(&user.ShortUsers{})
	return &su
}
