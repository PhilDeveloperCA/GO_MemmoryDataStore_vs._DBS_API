package Users

import (
	"net/http"

	_ "github.com/gorilla/mux"
)

type User struct {
	id       int32  `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Node struct {
	//next Node
}

type BinarySearchTree struct {
}

type BSTIndex struct {
	primary_key map[int]int
	bst_index   BinarySearchTree
}

type UserStore struct {
	users        []User
	userNameHash map[string]bool
	id_index     BSTIndex
}

type addUserReturn struct {
	user User
	err  error
}

// not valid syntax to initiate a Slice var Users = []User
func (store *UserStore) addUser(user User) addUserReturn {

	// how to append a slice -> current array
	store.users = append(store.users, user)
	//append(store.users, user)
	return addUserReturn{user, nil}
}

//

type UserSignup struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type UserResponse struct {
	JWT      string `json:"JWT"`
	Username string `json:"Username"`
}

// associate by pointer vs ????
func (*User) signUp() {

}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Refresh(w http.ResponseWriter, r *http.Request) {

}

func sample() {
	//valid syntax for make
	// Users2 := make([]User, 0, 1000000);s
}
