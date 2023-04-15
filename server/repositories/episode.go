package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type EpisodeRepository interface {
	FindEpisodesByMovie(MovieID int) ([]models.Episode, error)
	GetEpisodeByMovie(MovieID int, EpsodeID int) (models.Episode, error)
	GetEpisode(ID int) (models.Episode, error)
	CreateEpisode(Episode models.Episode) (models.Episode, error)
	UpdateEpisode(Episode models.Episode) (models.Episode, error)
	DeleteEpisode(Episode models.Episode, ID int) (models.Episode, error)
}

func RepositoryEpisode(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindEpisodesByMovie(movieID int) ([]models.Episode, error) {
	var episodes []models.Episode
	err := r.db.Preload("Movie").Find(&episodes).Error

	return episodes, err
}

func (r *repository) GetEpisode(ID int) (models.Episode, error) {
	var Episode models.Episode
	err := r.db.Preload("Movie").First(&Episode, ID).Error

	return Episode, err
}

func (r *repository) CreateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Preload("Movie").Create(&episode).Error

	return episode, err
}

func (r *repository) UpdateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Save(&episode).Error

	return episode, err
}

func (r *repository) DeleteEpisode(episode models.Episode, ID int) (models.Episode, error) {
	err := r.db.Delete(&episode, ID).Scan(&episode).Error

	return episode, err
}

func (r *repository) GetEpisodeByMovie(movieID int, episodeID int) (models.Episode, error) {
	var episode models.Episode
	err := r.db.Preload("Movie").Where("movie_id = ? AND episode_id = ?", movieID, episodeID).First(&episode).Error

	return episode, err
}
