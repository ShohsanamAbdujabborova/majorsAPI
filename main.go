package main

import (
	"API/controllers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server working")

	http.HandleFunc("/majors", controllers.MajorsController)

	http.ListenAndServe(":8080", nil)
}
