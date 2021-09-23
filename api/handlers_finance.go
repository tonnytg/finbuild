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

	// test
	database.DB.Table("exchanges").Create(&f)

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
	w.Write(b)
}
