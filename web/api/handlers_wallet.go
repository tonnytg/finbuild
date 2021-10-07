package api

import (
	"finbuild/pkg/db"
	"github.com/google/uuid"
	"net/http"
)

func getWallet(w http.ResponseWriter, r *http.Request) {

	// read db return and parse to struct
	q := r.URL.Query()
	walletID := q["id"]
	wl := db.GetWalletsByID(uuid.MustParse(walletID[0]))

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"wallets": wl,
	}

	JParse(w, mp1)
}
