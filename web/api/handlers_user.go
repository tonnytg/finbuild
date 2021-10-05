package api

import (
	"encoding/json"
	"finbuild/entity/finance"
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

	// new wallet
	wallet := finance.NewWallet(user.UserID)

	// save in database
	db.UserRegistry(user, wallet)

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

	// get only one user of slice query id
	user := db.GetUser(uuid.MustParse(userid[0]))

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"exchange": user,
	}

	JParse(w, mp1)
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	// read db return and parse to struct
	var wallet *finance.ShortWallet
	q := r.URL.Query()
	userid := q["id"]
	wallet = db.GetWallets(wallet, uuid.MustParse(userid[0]))

	var msg []map[string]interface{}

	f := finance.ShortWallet{
		WalletID: wallet.WalletID,
		Balance:  wallet.Balance,
	}

	// get an essential information form user
	su := db.GetShortUser(uuid.MustParse(userid[0]))

	mp1 := map[string]interface{}{
		"user": su,
	}
	msg = append(msg, mp1)

	mp2 := map[string]interface{}{
		"finance": f,
	}
	msg = append(msg, mp2)

	JParse(w, mp1, mp2)
}
