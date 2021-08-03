package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Username string
	Lastname string
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// responseString := "Hello World"
		data, err := json.Marshal(Person{"Albert", "Bellinger"})
		if err != nil {
			fmt.Println("Error")
		}
		w.Header().Add("content-type", "application/json")
		// w.Write([]byte(responseString))
		w.Write(data)
		return
		//
	})
	// mux.ErrMethodMismatch()
}
