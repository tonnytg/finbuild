package webserver

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tmpl = template.Must(template.ParseGlob("web/webserver/templates/*.html"))
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	exchanges := GetExchangesInfo("5f3ac43e-796a-4f89-993a-739124210d87")

	err := tmpl.ExecuteTemplate(w, "Index", exchanges)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}
