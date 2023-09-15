package hx

import (
	"github.com/labstack/echo/v4"
)

type (
	Ctx struct {
		echo.Context
	}
)

func Wrap(ctx echo.Context) *Ctx {
	return &Ctx{Context: ctx}
}

func Cast(ctx echo.Context) *Ctx {
	if _, ok := ctx.(*Ctx); !ok {
		ctx = Wrap(ctx)
	}
	return ctx.(*Ctx)
}

func (c *Ctx) Method() string {
	return c.Request().Method
}

func (c *Ctx) Path() string {
	return c.Request().URL.Path
}

// middleware for echo that wraps the context
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return next(Wrap(ctx))
	}
}
