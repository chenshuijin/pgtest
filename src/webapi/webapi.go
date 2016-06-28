//main.go
package main

import (
	"apis/root"
	"apis/users"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("init...")
	r := mux.NewRouter()
	// the request header content-type must be application/(text|json)
	r.Headers("Content-Type", "application/(text|json)")
	r.HandleFunc("/", root.DefaultRoute).Methods("GET")
	r.HandleFunc("/init", root.InitDatabase).Methods("GET")
	// Routes consist of a path and a handler function.
	r.HandleFunc("/users", users.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", users.CreateUser).Methods("POST")
	r.HandleFunc("/users/{uid}/relationships", users.GetRelationships).Methods("GET")
	r.HandleFunc("/users/{uid}/relationships/{ouid}", users.CreateRelationships).Methods("PUT")

	// Bind to a port and pass router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
