package handlers

import (
	episodeDto "dumbflix/dto/episode"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerEpisode struct {
	EpisodeRepository repositories.EpisodeRepository
}

func HandlerEpisode(EpisodeRepository repositories.EpisodeRepository) *handlerEpisode {
	return &handlerEpisode{EpisodeRepository}
}

func (h *handlerEpisode) FindEpisodesByMovie(c echo.Context) error {
	movieID, _ := strconv.Atoi(c.Param("movieID"))
	episodes, err := h.EpisodeRepository.FindEpisodesByMovie(movieID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	for i, m := range episodes {
		episodes[i].Thumbnail = path_file + m.Thumbnail
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episodes})
}

func (h *handlerEpisode) GetEpisodeByMovie(c echo.Context) error {
	movieID, _ := strconv.Atoi(c.Param("movieID"))
	episodeID, _ := strconv.Atoi(c.Param("episodeID"))

	episode, err := h.EpisodeRepository.GetEpisodeByMovie(movieID, episodeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	episode.Thumbnail = path_file + episode.Thumbnail

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episode})
}

func (h *handlerEpisode) GetEpisode(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	episode.Thumbnail = path_file + episode.Thumbnail

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episode})
}

func (h *handlerEpisode) CreateEpisode(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	movieID, _ := strconv.Atoi(c.FormValue("movie_id"))

	request := episodeDto.CreateEpisodeRequest{
		Title:     c.FormValue("title"),
		Thumbnail: dataFile,
		VideoLink: c.FormValue("video_link"),
		MovieID:   movieID,
	}

	fmt.Println(request)

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	episode := models.Episode{
		Title:     request.Title,
		Thumbnail: request.Thumbnail,
		VideoLink: request.VideoLink,
		MovieID:   request.MovieID,
	}

	data, err := h.EpisodeRepository.CreateEpisode(episode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)})
}

func (h *handlerEpisode) UpdateEpisode(c echo.Context) error {
	request := new(episodeDto.UpdateEpisodeRequest)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		episode.Title = request.Title
	}

	if request.Thumbnail != "" {
		episode.Thumbnail = request.Thumbnail
	}

	if request.VideoLink != "" {
		episode.VideoLink = request.VideoLink
	}

	if request.MovieID != 0 {
		if _, err := strconv.Atoi(strconv.Itoa(request.MovieID)); err == nil {
			episode.MovieID = request.MovieID
		}
	}

	data, err := h.EpisodeRepository.UpdateEpisode(episode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)})
}

func (h *handlerEpisode) DeleteEpisode(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.EpisodeRepository.DeleteEpisode(episode, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)})
}

func convertResponseEpisode(u models.Episode) models.EpisodeResponse {
	return models.EpisodeResponse{
		ID:        u.ID,
		Title:     u.Title,
		Thumbnail: u.Thumbnail,
		VideoLink: u.VideoLink,
		Movie:     u.Movie,
	}
}
