package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Task = "Denis"

type RequestBody struct {
	Task string `json:"task"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&RequestBody{})
	if err := json.NewEncoder(w).Encode(Task); err != nil {
		fmt.Println(err)
	}
	Task = RequestBody{}.Task
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,"+Task)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/Get", HelloHandler).Methods("GET")
	router.HandleFunc("/Post", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
git