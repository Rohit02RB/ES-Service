package main

import (
	"fmt"
	"net/http"

	"ES-Service/usecases"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("entered into ")
	router := mux.NewRouter()
	router.HandleFunc("/searchResult", usecases.FinalResult)
	// router.HandleFunc("/autocomplete",usecases.AutoComplete)
	http.ListenAndServe(":8082", router)
}
