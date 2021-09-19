package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"finbuild/entity/user"
)

// getUser return api response with user info
func getUser(w http.ResponseWriter, r *http.Request) {

	var msg []map[string]interface{}

	u := user.User{
		ID:         "1",
		FistName:   "Tonnytg",
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
		"message": "success",
		"user":    u,
	}
	msg = append(msg, mp1)

	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
