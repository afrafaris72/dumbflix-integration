package handlers

import (
	moviesDto "dumbflix/dto/movie"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerMovie struct {
	MovieRepository repositories.MovieRepository
}

func HandlerMovie(MovieRepository repositories.MovieRepository) *handlerMovie {
	return &handlerMovie{MovieRepository}
}

var path_file = "http://localhost:5000/uploads/"

func (h *handlerMovie) FindMovies(c echo.Context) error {
	movies, err := h.MovieRepository.FindMovies()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	for i, m := range movies {
		movies[i].Thumbnail = path_file + m.Thumbnail
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: movies})
}

func (h *handlerMovie) GetMovie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	movie, err := h.MovieRepository.GetMovie(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	movie.Thumbnail = path_file + movie.Thumbnail

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: movie})
}

func (h *handlerMovie) CreateMovie(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	categoryID, _ := strconv.Atoi(c.FormValue("category_id"))

	request := moviesDto.CreateMovieRequest{
		Title:       c.FormValue("title"),
		Thumbnail:   dataFile,
		Year:        c.FormValue("year"),
		CategoryID:  categoryID,
		Description: c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	movie := models.Movie{
		Title:       request.Title,
		Thumbnail:   request.Thumbnail,
		Year:        request.Year,
		CategoryID:  request.CategoryID,
		Description: request.Description,
	}

	data, err := h.MovieRepository.CreateMovie(movie)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseMovie(data)})
}

func (h *handlerMovie) UpdateMovie(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	categoryID, _ := strconv.Atoi(c.FormValue("category_id"))

	request := moviesDto.CreateMovieRequest{
		Title:       c.FormValue("title"),
		Thumbnail:   dataFile,
		Year:        c.FormValue("year"),
		CategoryID:  categoryID,
		Description: c.FormValue("description"),
	}

	id, _ := strconv.Atoi(c.Param("id"))

	movie, err := h.MovieRepository.GetMovie(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		movie.Title = request.Title
	}

	if request.Thumbnail != "" {
		movie.Thumbnail = request.Thumbnail
	}

	if request.Year != "" {
		movie.Year = request.Year
	}

	if request.CategoryID != 0 {
		if _, err := strconv.Atoi(strconv.Itoa(request.CategoryID)); err == nil {
			movie.CategoryID = request.CategoryID
		}
	}

	if request.Description != "" {
		movie.Description = request.Description
	}

	data, err := h.MovieRepository.UpdateMovie(movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseMovie(data)})
}

func (h *handlerMovie) DeleteMovie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	movie, err := h.MovieRepository.GetMovie(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.MovieRepository.DeleteMovie(movie, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseMovie(data)})
}

func convertResponseMovie(u models.Movie) models.MovieResponse {
	return models.MovieResponse{
		ID:          u.ID,
		Title:       u.Title,
		Thumbnail:   u.Thumbnail,
		Year:        u.Year,
		Category:    u.Category,
		Description: u.Description,
	}
}
