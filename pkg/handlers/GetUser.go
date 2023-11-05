package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nazhard/go-simple-crud/pkg/rawrr"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock users
	for _, user := range rawrr.Users {
		if user.Id == id {
			// If ids are equal send user as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			break
		}
	}
}