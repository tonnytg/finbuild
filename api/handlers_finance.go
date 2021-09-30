package api

import (
	"encoding/json"
	"finbuild/database"
	"finbuild/entity/finance"
	log "finbuild/pkg/logging"
	"fmt"
	"io/ioutil"
	"net/http"
)

// postExchange store in database action of exchange
func postExchange(w http.ResponseWriter, r *http.Request) {

	// read body response
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("problem to read body: %v", err)
		log.Msg("CRITICAL", msg)

	}

	// convert body response to struct
	var f finance.Asset
	err = json.Unmarshal(body, &f)
	if err != nil {
		msg := fmt.Sprintf("problem to unmarshal body: %v", err)
		log.Msg("ERROR", msg)
	}

	// save database
	database.DB.Table("exchanges").Create(&f)

	// update Balance
	a := finance.Account{}
	a.UpdateBalance(f.UserID, f.Action, f.Quantity, f.Price)
	a.RegisterAsset(f.UserID, f.Action, f.ID, f.Quantity, f.Price, f.Date)

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"exchange": f,
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
	_, err = w.Write(b)
	if err != nil {
		return
	}
}
