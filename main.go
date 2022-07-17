package main

import (
	"go-api-student-mysql-gorm/router"
	"net/http"
)

func main() {
	r := router.Router()
	err := http.ListenAndServe(":9000", r)
	if err != nil {
		panic(err)
	}
}


