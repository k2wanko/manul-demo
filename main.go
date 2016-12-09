package main

import (
	"log"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handleIndex)
	log.Fatal(e.Start(":8080"))
}

func handleIndex(c echo.Context) error {
	return c.String(200, "ok")
}
