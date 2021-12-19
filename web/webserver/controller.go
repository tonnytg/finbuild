package webserver

import (
	"finbuild/entity/exchange"
	"finbuild/entity/user"
	"finbuild/pkg/db"
	"github.com/google/uuid"
)

func GetUserInfo(s string) *user.User {
	user := db.GetUser(uuid.MustParse(s))
	return user
}

func GetExchangesInfo(s string) []exchange.Exchanges {
	exchanges := db.GetExchanges(uuid.MustParse(s))
	return exchanges
}

