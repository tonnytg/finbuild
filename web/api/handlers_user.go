package api

import (
	"encoding/json"
	entity "finbuild/entity/finance"
	"finbuild/pkg/db"
	"github.com/google/uuid"
	"net/http"

	"finbuild/entity/user"
)

func registerUser(w http.ResponseWriter, r *http.Request) {

	// new user
	user := user.NewUser()

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&user)
	db.DB.Model(&user).Create(&user)

	// new wallet
	wallet := entity.NewWallet(user.UserID)
	db.DB.Table("wallets").Model(&entity.Wallet{}).Create(&wallet)

	type RegistryReturn struct {
		WalletID uuid.UUID
		UserID   uuid.UUID
	}

	rr := RegistryReturn{
		wallet.WalletID,
		user.UserID,
	}

	mp1 := map[string]interface{}{
		"registry": rr,
	}
	JParse(w, mp1)
}

// getUser return api response with user info
func getUser(w http.ResponseWriter, r *http.Request) {

	// read db return and parse to struct
	q := r.URL.Query()
	userid := q["id"]

	dbuser := &user.User{}
	db.DB.First(&dbuser, userid).Scan(&dbuser)

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"exchange": dbuser,
	}

	JParse(w, mp1)
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	// get balance
	type Account struct {
		WalletID string
		Balance  float64
	}

	var account Account
	// read db return and parse to struct
	q := r.URL.Query()
	userid := q["id"]
	db.DB.Table("wallets").First(&account, "user_id = ?", userid).Scan(&account).Find(&Account{})

	var msg []map[string]interface{}

	f := Account{
		WalletID: account.WalletID,
		Balance:  account.Balance,
	}

	type Users struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
	}
	u := Users{}
	db.DB.Table("users").First(&u, "user_id = ?", userid).Scan(&u).Find(&Users{})

	mp1 := map[string]interface{}{
		"user": u,
	}
	msg = append(msg, mp1)

	mp2 := map[string]interface{}{
		"finance": f,
	}
	msg = append(msg, mp2)

	JParse(w, mp1, mp2)
}
