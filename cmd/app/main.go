package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pharrisee/base/content"
	"github.com/pharrisee/base/internal/hx"
)

func main() {
	e := echo.New()
	e.Renderer = content.Renderer()
	e.Use(hx.Middleware) // MUST be the first middleware used, to provide to other middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(content.StaticFS),
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
