package main

import (
	"os"

	"github.com/go-travelpayouts/travelr"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	gr := e.Group("/api/v1/public")
	gr.POST("/signature", travelr.GetSignatureFromRequest)

	err := e.Start(":" + getPort())
	if err != nil {
		panic(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	return port
}
