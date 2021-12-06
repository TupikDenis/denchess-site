package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TupikDenis/denchess-site.git/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock books
	for _, user := range mocks.Users {
		if user.Id == id {
			// If ids are equal send book as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			break
		}
	}
}
