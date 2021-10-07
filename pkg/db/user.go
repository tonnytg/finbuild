package db

import (
	"finbuild/entity/user"
	"finbuild/entity/wallet"
	"github.com/google/uuid"
)

// UserRegistry get user and wallet information to sabe in database
func UserRegistry(u *user.User, w *wallet.Wallet) {

	DB.Model(&u).Create(&u)

	DB.Table("wallets").Model(&wallet.Wallet{}).Create(&w)
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
