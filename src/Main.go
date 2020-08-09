package main

import (
	"encoding/json"
	"goTest/src/model"
	"log"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	nicolas := model.CreateUser("nicolas", "canicatti", 58)

	response, _ := json.Marshal(*nicolas)
	w.Write(response)
}

//func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte(`{"message": "hello world"}`))
//
//
//}

func main() {

	http.HandleFunc("/user", getUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
