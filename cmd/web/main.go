package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"np_finance/internal/log"
	"np_finance/internal/run_handle"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.gohtml")),
	}
}

type Handler interface {
	SetupRoutes(app *echo.Echo) error
}

var Handlers []Handler

func main() {
	app := echo.New()
	//app.Use(middleware.Logger())
	log.NewLogger()
	app.Use(log.LoggingMiddleware)
	app.Use(middleware.Recover())
	app.Static("/images", "web/images")
	app.Static("/css", "web/css")
	app.Renderer = newTemplates()

	Handlers = append(Handlers, &run_handle.RunHandlers{})

	for _, handler := range Handlers {
		err := handler.SetupRoutes(app)
		if err != nil {
			log.Logger.LogFatal().Msgf("Error: %v \nin Handler: %v", err, handler)
		}
	}

	app.GET("/", func(c echo.Context) error {
		return c.Redirect(302, "/run")
	})

	app.Logger.Fatal(app.Start(":8080"))

}
