package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func testHandle(c echo.Context) error {
	return c.String(http.StatusOK, "You found me, now do something else")
}

func main() {
	vergil := echo.New()

	vergil.Use(middleware.Logger())
	vergil.Use(middleware.Recover())
	vergil.Static("/static", "static")

	vergil.GET("/", testHandle)

	vergil.Logger.Fatal(vergil.Start(":8900"))
}
