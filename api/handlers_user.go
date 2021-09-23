package api

import (
	"encoding/json"
	"finbuild/database"
	"fmt"
	"net/http"

	"finbuild/entity/user"
)

// getUser return api response with user info
func getUser(w http.ResponseWriter, r *http.Request) {

	// read database return and parse to struct
	dbuser := &user.User{}
	database.DB.First(&dbuser, "1").Scan(&dbuser)

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
