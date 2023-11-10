package Repository

type UserRepository struct {
	// Add database connection properties here
}

func NewChecklistRepository() *UserRepository {
	return &UserRepository{
		// Initialize database connection
	}
}

func (r *UserRepository) ListItems() []Item {
	// Database logic to retrieve items
	// ...

	return nil // Replace with actual data retrieval logic
}

type Item struct {
	name  string
	model bool
}
