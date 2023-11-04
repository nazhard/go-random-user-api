package main

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

// User struct (Model)
type User struct {
  Name  string `json:"name"`
  Phone string `json:"phone"`
  Email string `json:"email"`
  Age   int8   `json:"age"`
}

// Init users var as a slice User struct
var users []User

// Get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(users)
}

// Get single user
func getUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r) // Gets params
  // Looping through users and find one with the id from the params
  for _, item := range users {
    if item.Name == params["name"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&User{})
}

// Add new user
func createUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var user User
  _ = json.NewDecoder(r.Body).Decode(&user)
  users = append(users, user)
  json.NewEncoder(w).Encode(user)
}

// Update user
func updateUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for idx, item := range users {
    if item.Name == params["name"] {
      users = append(users[:idx], users[idx+1:]...)
      var user User
      _ = json.NewDecoder(r.Body).Decode(&user)
      user.Name = params["name"]
      users = append(users, user)
      json.NewEncoder(w).Encode(user)
      return
    }
  }
}

// Delete user
func deleteUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for idx, item := range users {
    if item.Name == params["name"] {
      users = append(users[:idx], users[idx+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(users)
}

// Main function
func main() {
  // Init router
  r := mux.NewRouter()

  // Bruh, bad data
  users = append(users, User{Name: "Agus", Phone: "62xxx-xxxx-xxxx", Email: "Agus@example.com", Age: 19})
  users = append(users, User{Name: "Fadhil", Phone: "62xxx-xxxx-xxxx", Email: "twins@example.com", Age: 23})
  users = append(users, User{Name: "Wawan", Phone: "62xxx-xxxx-xxxx", Email: "Wawan@example.com", Age: 29})
  users = append(users, User{Name: "Agan", Phone: "62xxx-xxxx-xxxx", Email: "krocofb@example.com", Age: 18})
  users = append(users, User{Name: "Fadlan", Phone: "62xxx-xxxx-xxxx", Email: "twins@example.com", Age: 23})
  users = append(users, User{Name: "Mansur", Phone: "62xxx-xxxx-xxxx", Email: "trillionare@example.com", Age: 45})
  users = append(users, User{Name: "Luwak", Phone: "62xxx-xxxx-xxxx", Email: "kopi@example.com", Age: 96})
  users = append(users, User{Name: "Asep Bensin", Phone: "62xxx-xxxx-xxxx", Email: "asepbumi@example.com", Age: 19})
  users = append(users, User{Name: "Samsul Arif", Phone: "62xxx-xxxx-xxxx", Email: "samsul@example.com", Age: 34})
  users = append(users, User{Name: "Nazhard", Phone: "62xxx-xxxx-xxxx", Email: "asli.masih.pemula@example.com", Age: 15})
  
  // Route handles & endpoints
  r.HandleFunc("/users", getUsers).Methods("GET")
  r.HandleFunc("/users/{name}", getUser).Methods("GET")
  r.HandleFunc("/users", createUser).Methods("POST")
  r.HandleFunc("/users/{name}", updateUser).Methods("PUT")
  r.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")

  // Start server
  log.Fatal(http.ListenAndServe(":3000", r))
}