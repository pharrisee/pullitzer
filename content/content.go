package content

import (
	"embed"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/unrolled/render"
)

//go:embed static
var StaticFS embed.FS

//go:embed views
var ViewsFS embed.FS

type (
	// Renderer is a custom renderer for echo
	Views struct {
		r *render.Render
	}
)

func Renderer() *Views {
	opts := render.Options{
		FileSystem: &render.EmbedFileSystem{FS: ViewsFS},
	}
	return &Views{r: render.New(opts)}
}

func (v *Views) Render(w io.Writer, name string, data any, ctx echo.Context) error {
	// use 0 as echo will fill in the correct status code later
	return v.r.HTML(w, 0, name, data)
}
