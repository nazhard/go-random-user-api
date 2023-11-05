package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nazhard/go-simple-crud/pkg/rawrr"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rawrr.Users)
}