package Controller

import (
	"awesomeProject/Service"
	"net/http"
)

type ChecklistController struct {
	service           *Service.ChecklistService
	HomeScreenHandler http.HandlerFunc
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

func NewChecklistController(service *Service.ChecklistService) *ChecklistController {
	return &ChecklistController{service: service}
}

func (c *ChecklistController) HomeHandler(w http.ResponseWriter, r *http.Request) []Service.ProcessedItem {
	return c.service.ListItems()
}
