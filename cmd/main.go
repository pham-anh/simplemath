package main

import (
	"math/rand"
	"time"

	"simplemath/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Seeded RNG for generation (inject into handler)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	h := handler.NewSubmitHandler(r)

	e.POST("/", h.HandleSubmit)
	e.GET("/", func(c echo.Context) error { return c.File("statics/index.html") })

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
