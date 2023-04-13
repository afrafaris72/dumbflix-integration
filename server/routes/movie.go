package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/labstack/echo/v4"
)

func MovieRoutes(e *echo.Group) {
	MovieRepository := repositories.RepositoryMovie(mysql.DB)

	h := handlers.HandlerMovie(MovieRepository)

	e.GET("/movies", h.FindMovies)
	e.GET("/movie/:id", middleware.Auth(h.GetMovie))
	e.POST("/movie", middleware.Auth(middleware.UploadFile(h.CreateMovie)))
	e.PATCH("/movie/:id", middleware.Auth(middleware.UploadFile(h.UpdateMovie)))
	e.DELETE("/movie/:id", middleware.Auth(h.DeleteMovie))
}
