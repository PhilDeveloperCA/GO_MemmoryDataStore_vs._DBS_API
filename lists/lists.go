package lists

import (
	_ "github.com/gorilla/mux"
)

type List struct {
	id          int
	store       int
	creator_id  int
	created_at  string
	date        string
	group_users []int
}

type ListStore struct {
	lists []List
}

func init() {

}
