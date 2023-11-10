// main.go
package main

import (
	"awesomeProject/Controller"
	"awesomeProject/Repository"
	"awesomeProject/Service"
	"net/http"
)

func main() {
	repo := Repository.NewChecklistRepository()
	service := Service.NewChecklistService(repo)
	controller := Controller.NewChecklistController(service)

	http.HandleFunc("/home", controller.HomeHandler)
	http.HandleFunc("/", controller.HomeScreenHandler)
	// Setup other routes...

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Handle error
	}
}
