package run_handle

import "github.com/labstack/echo/v4"

type RunHandlers struct {
	// SetupRoutes(app *echo.Echo) error
}

func (h *RunHandlers) SetupRoutes(app *echo.Echo) error {
	app.GET("/run", func(c echo.Context) error {
		return c.String(200, "<hr>Hello, World! !!!")
	})
	return nil
}
