package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TupikDenis/denchess-site.git/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock Books
	for index, user := range mocks.Users {
		if user.Id == id {
			// Delete book and send response if the book Id matches dynamic Id
			mocks.Users = append(mocks.Users[:index], mocks.Users[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
