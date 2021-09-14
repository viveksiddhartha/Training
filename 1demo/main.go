package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"Training/1demo/datastore"
	"Training/1demo/handlers"
)

const (
	WEBPORT = ":8080"
)

func main() {
	fmt.Println("My Application is starting")

	datastore.DBCONN()

	router := mux.NewRouter()

	http.Handle("/", router)

	router.HandleFunc("/home", homefunc)
	router.HandleFunc("/createuser", handlers.CreateUser)
	router.HandleFunc("/getUser", handlers.GetUsers)
	router.HandleFunc("/UpdateUser", handlers.UpdateUsers)
	router.HandleFunc("/DeleteUser", handlers.DeleteUsers)

	http.ListenAndServe(WEBPORT, nil)

}

func homefunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "We have received request")

}
