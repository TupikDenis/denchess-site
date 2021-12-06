package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/TupikDenis/denchess-site.git/pkg/mocks"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	// Append to the Book mocks
	user.Id = rand.Intn(100)
	mocks.Users = append(mocks.Users, user)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
