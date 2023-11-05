package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nazhard/go-simple-crud/pkg/rawrr"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the rawrr Users
	for index, user := range rawrr.Users {
		if user.Id == id {
			// Delete user and send response if the user Id matches dynamic Id
			rawrr.Users = append(rawrr.Users[:index], rawrr.Users[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}