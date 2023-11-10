package main

import (
	"awesomeProject/Controller"
	"awesomeProject/Facade"
	"net/http"
)

func main() {
	// Initialize the Facade and pass it to the Controller
	facade := Facade.NewChecklistFacade()

	// Initialize the Controller with the Facade
	checklistController := Controller.NewChecklistController(facade)

	// Register the handler
	http.HandleFunc("/home", checklistController.HomeHandler)

	http.HandleFunc("/", checklistController.HomeScreenHandler)

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
