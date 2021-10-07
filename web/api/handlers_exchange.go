package api

import (
	"encoding/json"
	"finbuild/entity/exchange"
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
	var e exchange.Exchanges
	err = json.Unmarshal(body, &e)
	if err != nil {
		msg := fmt.Sprintf("problem to unmarshal body: %v", err)
		log.Msg("ERROR", msg)
	}

	// update Balance
	db.UpdateBalance(&e)

	// save exchange
	db.RegisterExchange(&e)

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"exchange": e,
	}

	JParse(w, mp1)
}
