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

	var msg []map[string]interface{}

	u := user.User{
		ID:         "1",
		FirstName:  "Tonnytg",
		LastName:   "TG",
		SocialID:   "001.001.001-01",
		Age:        50,
		Email:      "tonnytg@domain.com",
		ValidEmail: true,
		Phone:      "+5522999999999",
		Sex:        "male",
		Sign:       "Aquarius",
		Address:    "São Paulo, Brazil",
	}

	mp1 := map[string]interface{}{
		"user": u,
	}
	msg = append(msg, mp1)

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

func getUserDB(w http.ResponseWriter, r *http.Request) {

	//id := 1
	dbuser := &user.User{}
	database.DB.First(&dbuser, "1").Scan(&dbuser)
	fmt.Println(dbuser)
	b, _ := json.Marshal(dbuser)
	w.Write(b)
}
