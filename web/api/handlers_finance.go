package api

import (
	"encoding/json"
	"finbuild/entity/finance"
	"finbuild/pkg/db"
	log "finbuild/pkg/logging"
	"fmt"
	"github.com/google/uuid"
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
	db.UpdateBalance(&f)

	// save exchange
	db.RegisterExchange(&f)

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
	wallet := db.GetWalletsByID(uuid.MustParse(walletID[0]))

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"wallets": wallet,
	}

	JParse(w, mp1)
}
