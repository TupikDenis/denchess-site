package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/TupikDenis/denchess-site.git/pkg/mocks"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedUser models.User
	json.Unmarshal(body, &updatedUser)

	// Iterate over all the mock Books
	for index, user := range mocks.Users {
		if user.Id == id {
			// Update and send response when book Id matches dynamic Id
			user.Email = updatedUser.Email
			user.Password = updatedUser.Password
			user.LastName = updatedUser.LastName
			user.FirstName = updatedUser.FirstName

			mocks.Users[index] = user

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
