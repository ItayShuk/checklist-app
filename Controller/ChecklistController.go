package Controller

import (
	"awesomeProject/Facade"
	"net/http"
)

type ChecklistController struct {
	facade            Facade.ChecklistFacade
	HomeHandler       http.HandlerFunc
	HomeScreenHandler http.HandlerFunc
}

func NewChecklistController(f Facade.ChecklistFacade) *ChecklistController {
	return &ChecklistController{
		facade:            f,
		HomeHandler:       homeHandler(), // Assign the homeHandler function
		HomeScreenHandler: homeScreenHandler(),
	}
}

func homeScreenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ... use facade to get data and write to response ...
		_, err := w.Write([]byte("Hey, here is your checklist"))
		if err != nil {
			return
		}
	}
}

func homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ... use facade to get data and write to response ...
		_, err := w.Write([]byte("Home Page"))
		if err != nil {
			return
		}
	}
}
