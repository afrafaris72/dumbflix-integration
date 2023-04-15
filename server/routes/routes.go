package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	MovieRoutes(e)
	CategoryRoutes(e)
	TransactionRoutes(e)
	EpisodeRoutes(e)
	AuthRoutes(e)
}
