package guide

import "github.com/muathendirangu/lavida-api/category"

//Guide of a trip
type Guide struct {
	ID       int64             `json:"tour_guide_id"`
	Name     string            `json:"tour_guide_name"`
	Email    string            `json:"tour_guide_email"`
	Number   string            `json:"tour_guide_number"`
	Category category.Category `json:"category_id"`
}

//Repository defines how to store Guide
type Repository interface {
	Store(guide *Guide) error
}