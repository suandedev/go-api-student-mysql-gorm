package middleware

import (
	"encoding/json"
	"go-api-student-mysql-gorm/functions"
	"go-api-student-mysql-gorm/model"
	"net/http"
)

// func create new user
func Create(w http.ResponseWriter, r *http.Request) {
	// create
	rs := functions.Create("suande", "asd@gmail.com", 22)
	// db.Create(&model.User{Name: name, Email: email, Age: age})
	if rs.Error != nil {
		panic(rs.Error)
	}
	
	var usr model.User
	rs.Scan(&usr)
	json.NewEncoder(w).Encode(usr)
}

// func read all users
func ReadAll(w http.ResponseWriter, r *http.Request) {
	rs := functions.ReadAll()
	// read all
	var usr []model.User
	rs.Scan(&usr)
	json.NewEncoder(w).Encode(usr)
}

// func update user
func Update(w http.ResponseWriter, r *http.Request) {
	rs := functions.Update(4, "susuande", "ms@gmail.com", 23)
	// update
	var usr model.User
	rs.Scan(&usr)
	json.NewEncoder(w).Encode(usr)
}

// // func delete user
func Delete(w http.ResponseWriter, r *http.Request) {
	// delete
	functions.Delete(3)
	
	json.NewEncoder(w).Encode("berhasil")
}

// // func read user by id
func ReadById(w http.ResponseWriter, r *http.Request) {
	rs := functions.ReadById(4)
	var usr model.User
	rs.Scan(&usr)
	json.NewEncoder(w).Encode(usr)
}