package user

import (
	"encoding/json"
	"net/http"

	"github.com/TupikDenis/denchess-site.git/pkg/mocks"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Users)
}
