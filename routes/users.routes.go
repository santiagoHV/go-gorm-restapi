package routes

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users Handler"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User Handler"))
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Users Handler"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Users Handler"))
}
