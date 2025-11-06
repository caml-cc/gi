package gi

import (
	"gitignore/internal/utils"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Generate(w http.ResponseWriter, r *http.Request) {
	chosen := mux.Vars(r)["catchall"]

	if chosen == "list" {
		listing, err := utils.ListDir("templates")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(listing)
		return
	} else {
		template := TemplateMaker(strings.Split(chosen, ","))
		if template == nil {
			http.Error(w, "", 404)
			return
		}
		w.Write(template)
		return
	}
}
