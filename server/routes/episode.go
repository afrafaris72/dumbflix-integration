package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/labstack/echo/v4"
)

func EpisodeRoutes(e *echo.Group) {
	EpisodeRepository := repositories.RepositoryEpisode(mysql.DB)

	h := handlers.HandlerEpisode(EpisodeRepository)

	e.GET("/episodes", middleware.Auth(h.FindEpisodesByMovie))
	e.GET("/episode/:id", middleware.Auth(h.GetEpisode))
	e.POST("/episode", middleware.Auth(middleware.UploadFile(h.CreateEpisode)))
	e.PATCH("/episode/:id", middleware.Auth(middleware.UploadFile(h.UpdateEpisode)))
	e.DELETE("/episode/:id", middleware.Auth(h.DeleteEpisode))
}
