package utils

// item.go (or you can define it in an appropriate file)
type Item struct {
	Description string `bson:"description"`
	Done        bool   `bson:"done"`
}
