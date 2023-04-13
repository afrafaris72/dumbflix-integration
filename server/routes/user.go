package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	UserRepositroy := repositories.RepositoryUser(mysql.DB)

	h := handlers.HandlerUser(UserRepositroy)

	e.GET("/users", h.FindUsers)
	e.DELETE("/user/:id", h.DeleteUser)
}
