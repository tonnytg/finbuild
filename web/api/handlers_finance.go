package api

import (
	"encoding/json"
	"finbuild/entity/finance"
	"finbuild/pkg/db"
	log "finbuild/pkg/logging"
	"fmt"
	"io/ioutil"
	"net/http"
)

// postExchange store in db action of exchange
func postExchange(w http.ResponseWriter, r *http.Request) {

	// read body response
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("problem to read body: %v", err)
		log.Msg("CRITICAL", msg)
	}

	// convert body response to struct
	var f finance.Exchanges
	err = json.Unmarshal(body, &f)
	if err != nil {
		msg := fmt.Sprintf("problem to unmarshal body: %v", err)
		log.Msg("ERROR", msg)
	}

	// update Balance
	db.UpdateBalance(f.WalletID, f.Action, f.Quantity, f.Price, f.Tax, f.Date)
	// save exchange
	db.RegisterExchange(f.WalletID, f.AssetID, f.Action, f.Tax, f.Quantity, f.Price, f.Date)

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"exchange": f,
	}

	JParse(w, mp1)
}

func getWallet(w http.ResponseWriter, r *http.Request) {

	// read db return and parse to struct
	q := r.URL.Query()
	walletID := q["id"]

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"wallets": walletID,
	}

	JParse(w, mp1)
}
