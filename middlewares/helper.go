package middlewares

import (
	"github.com/caiomp87/go-barber/models"
	"github.com/gofiber/fiber/v2"
)

func getInterceptorService() map[string]string {
	return interceptedRoutes
}

func formatInterceptorValues(c *fiber.Ctx) string {
	return c.Method() + c.App().GetRoute(c.Route().Name).Path
}

func isPublic(interceptorServices map[string]string, interceptorValue string) bool {
	return interceptorServices[interceptorValue] == "public"
}

func BuildRoutes(app *fiber.App) []models.Route {
	routes := make([]models.Route, 0)

	for _, route := range app.GetRoutes(true) {
		routes = append(routes, models.Route{
			Method: route.Method,
			Name:   route.Name,
			Path:   route.Path,
			Params: route.Params,
		})
	}

	return routes
}
