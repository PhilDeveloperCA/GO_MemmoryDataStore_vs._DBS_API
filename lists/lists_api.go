package lists

import (
	"net/http"

	_ "github.com/gorilla/mux"
)

func getList(w http.ResponseWriter, r *http.Request) {
	//http.HandleFunc("/")
	var listid = r.URL.Path
}

func addList()

func deleteList()

func updateList()
