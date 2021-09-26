package api

import (
	"encoding/json"
	"finbuild/database"
	"finbuild/entity/finance"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// postExchange store in database action of exchange
func postExchange(w http.ResponseWriter, r *http.Request) {

	// read body response
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR: problem to read body", err)
	}

	// convert body response to struct
	var f finance.Asset
	err = json.Unmarshal(body, &f)
	if err != nil {
		log.Println("ERROR: problem to unmarshal body", err)
	}

	// save database
	database.DB.Table("exchanges").Create(&f)

	// update Balance
	a := finance.Account{}
	a.UpdateBalance(f.UserID, f.Action, f.Quantity, f.Price)
	a.RegisterAsset(f.UserID, f.Action, f.Quantity, f.Price)

	// create a slice of interface to receive json content
	mp1 := map[string]interface{}{
		"exchange": f,
	}

	if f.Action == "BUY" {
		fmt.Printf("user buy many %f of %s\n", f.Quantity, f.ID )
	}
	if f.Action == "SELL" {
		fmt.Printf("user sell many %f of %s\n", f.Quantity, f.ID )
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
