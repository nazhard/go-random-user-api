package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nazhard/go-simple-crud/pkg/rawrr"
	"github.com/nazhard/go-simple-crud/pkg/models"
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

	// Iterate over all the mock Users
	for index, user := range rawrr.Users {
		if user.Id == id {
			// Update and send response when user Id matches dynamic Id
			user.Name = updatedUser.Name
			user.Phone = updatedUser.Phone
			user.Email = updatedUser.Email
			user.Age = updatedUser.Age
			user.Student = updatedUser.Student

			rawrr.Users[index] = user

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}