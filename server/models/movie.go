package models

import "time"

type Movie struct {
	ID          int              `json:"id"`
	Title       string           `json:"title" gorm:"type VARCHAR(255)"`
	Thumbnail   string           `json:"thumbnail" gorm:"type VARCHAR(255)"`
	Year        string           `json:"year" gorm:"type VARCHAR(255)"`
	CategoryID  int              `json:"category_id"`
	Category    CategoryResponse `json:"category"`
	Description string           `json:"description" gorm:"type VARCHAR(255)"`
	CreatedAt   time.Time        `json:"-"`
	UpdatedAt   time.Time        `json:"-"`
}

type MovieResponse struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Thumbnail   string           `json:"thumbnail"`
	Year        string           `json:"year"`
	Category    CategoryResponse `json:"category"`
	Description string           `json:"description"`
}

func (MovieResponse) TableName() string {
	return "MovieResponses"
}
