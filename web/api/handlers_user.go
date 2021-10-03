package api

import (
	"encoding/json"
	entity "finbuild/entity/finance"
	"finbuild/pkg/db"
	"fmt"
	"github.com/google/uuid"
	"net/http"

	"finbuild/entity/user"
)

func registerUser(w http.ResponseWriter, r *http.Request) {

	// genId create a UUID for new user
	genID := uuid.New()

	decoder := json.NewDecoder(r.Body)
	var u user.User
	err := decoder.Decode(&u)

	u.UserID = genID
	db.DB.Model(&user.User{}).Create(&u)

	// create a blank wallet to new User
	var wallet entity.Wallet
	wallet.WalletID = uuid.New()
	wallet.UserID = u.UserID
	wallet.Balance = 0
	db.DB.Table("wallets").Model(&entity.Wallet{}).Create(&wallet)

	type RegistryReturn struct {
		WalletID uuid.UUID
		UserID uuid.UUID
	}

	rr :=  RegistryReturn{
		wallet.WalletID,
		u.UserID,
	}

	mp1 := map[string]interface{}{
		"registry": rr,
	}

	var msg []map[string]interface{}
	msg = append(msg, mp1)

	jSend := Response{
		Status:  "success",
		Data:    msg,
		Message: "registry with success",
	}

	b, err := json.Marshal(jSend)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
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

	// create a map for json template return
	var msg []map[string]interface{}
	msg = append(msg, mp1)

	// json template to return
	jSend := Response{
		Status:  "success",
		Data:    msg,
		Message: "test",
	}

	b, err := json.Marshal(jSend)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func GetRoot(w http.ResponseWriter, r *http.Request) {

	// get balance
	type Account struct {
		WalletID string
		Balance float64
	}

	var account Account
	// read db return and parse to struct
	q := r.URL.Query()
	userid := q["id"]
	db.DB.Table("wallets").First(&account, "user_id = ?", userid).Scan(&account).Find(&Account{})

	var msg []map[string]interface{}

	f := Account{
		WalletID: account.WalletID,
		Balance: account.Balance,
	}

	type Users struct {
		ID string `json:"id"`
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

	jSend := Response{
		Status:  "success",
		Data:    msg,
		Message: "test",
	}

	b, err := json.Marshal(jSend)
	if err != nil {
		fmt.Println("error:", err)

	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
