package api

import (
	"encoding/json"
	"finbuild/entity/wallet"

	"finbuild/pkg/db"
	"github.com/google/uuid"
	"net/http"

	"finbuild/entity/user"
)

func registerUser(w http.ResponseWriter, r *http.Request) {

	// new us
	us := user.NewUser()

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&us)

	// new wallet call be wl
	wl := wallet.NewWallet(us.UserID)

	// save in database
	db.UserRegistry(us, wl)

	type RegistryReturn struct {
		WalletID uuid.UUID
		UserID   uuid.UUID
	}

	rr := RegistryReturn{
		wl.WalletID,
		us.UserID,
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
	us := db.GetUser(uuid.MustParse(userid[0]))

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"exchange": us,
	}

	JParse(w, mp1)
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	// read db return and parse to struct
	var wl *wallet.ShortWallet
	q := r.URL.Query()
	userid := q["id"]
	wl = db.GetWalletsByUser(wl, uuid.MustParse(userid[0]))

	var msg []map[string]interface{}

	f := wallet.ShortWallet{
		WalletID: wl.WalletID,
		Balance:  wl.Balance,
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
