package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
 )

// Strucs
var people []Person

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

func main(){
  // Manually popullate a few People.
  people = append(people, Person{ID: "1", Firstname: "Tony", Lastname: "Chemex", Address: &Address{City: "Rome", State: "Rome"}})
  people = append(people, Person{ID: "2", Firstname: "Bobby", Lastname: "Giro", Address: &Address{City: "Manhattan", State: "NYC"}})
  people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Mc"})

  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}

// Routes
func GetPeople(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}
