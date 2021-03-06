package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/foundation-13/gpr/pkg/log"
	"github.com/foundation-13/gpr/pkg/api/content"
)

func main() {
	log.InitLog(true)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.L.Info("api started")
	m := content.NewManager()
	content.Assemble(e, m)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "green"})
	})

	e.Logger.Fatal(e.Start(":8765"))
}
