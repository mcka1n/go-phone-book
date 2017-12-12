package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
 )

// our main function
func main(){
  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeople).Methods("GET")

  log.Fatal(http.ListenAndServe(":8000", router))
}

// Routes

func GetPeople(w http.ResponseWriter, r *http.Request) {}
