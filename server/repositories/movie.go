package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	FindMovies() ([]models.Movie, error)
	GetMovie(ID int) (models.Movie, error)
	CreateMovie(movie models.Movie) (models.Movie, error)
	UpdateMovie(movie models.Movie) (models.Movie, error)
	DeleteMovie(movie models.Movie, ID int) (models.Movie, error)
}

func RepositoryMovie(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindMovies() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.db.Preload("Category").Find(&movies).Error

	return movies, err
}

func (r *repository) GetMovie(ID int) (models.Movie, error) {
	var movies models.Movie

	err := r.db.Preload("Category").First(&movies, ID).Error

	return movies, err
}

func (r *repository) CreateMovie(movie models.Movie) (models.Movie, error) {
	err := r.db.Preload("Category").Create(&movie).Error

	return movie, err
}

func (r *repository) UpdateMovie(movie models.Movie) (models.Movie, error) {
	err := r.db.Save(&movie).Error

	return movie, err
}

func (r *repository) DeleteMovie(movie models.Movie, ID int) (models.Movie, error) {
	err := r.db.Delete(&movie, ID).Scan(&movie).Error

	return movie, err
}
