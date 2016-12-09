package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	e.GET("/", handleIndex)
	s := standard.New(":8080")
	s.SetHandler(e)
	log.Fatal(s.ListenAndServe())
}

func handleIndex(c echo.Context) error {
	return c.String(200, "ok")
}
