package api

import (
	"finbuild/entity/user"
	"net/http"
)

// getUser return api response with user info
func getUser(w http.ResponseWriter, r *http.Request) {

	user, _ := user.GetUser()
	apiResponse(user, w)
}

