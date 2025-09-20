package main

import (
	"math/rand"
	"time"

	"simplemath/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())

	// Seeded RNG for generation (inject into handler)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	h := handler.NewSubmitHandler(r)

	// Apply the rate limiter middleware to the POST endpoint.
	// We'll limit it to 1 request per second from each IP address.
	ratelimit := middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(middleware.RateLimiterMemoryStoreConfig{
			Rate:      30,              // 60 requests
			Burst:     5,               // with a burst of 5
			ExpiresIn: 1 * time.Minute, // per minute
		}),
	})

	e.POST("/", h.HandleSubmit, ratelimit)

	e.GET("/", func(c echo.Context) error { return c.File("statics/index.html") }, ratelimit)
	e.File("/favicon.png", "statics/favicon.png")

	// Serve static CSS files
	e.Static("/css", "statics/css")

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
