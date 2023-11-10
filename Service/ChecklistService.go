package Service

import (
	"awesomeProject/Repository"
)

type ChecklistService struct {
	repo *Repository.UserRepository
}

func NewChecklistService(repo *Repository.UserRepository) *ChecklistService {
	return &ChecklistService{repo: repo}
}

func (s *ChecklistService) ListItems() []ProcessedItem {
	// Call the repository to get items
	items := s.repo.ListItems()
	itemsProcessed := make([]ProcessedItem, len(items))
	return itemsProcessed
}

type ProcessedItem struct {
	name string
	mode bool
}
