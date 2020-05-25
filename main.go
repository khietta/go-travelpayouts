package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type RoundTripRequest struct {
// 	Segments []travelr.SegmentRequest `json:"segments"`
// }

// func SearchRoundTrip(ctx echo.Context) error {
// 	var objRequest RoundTripRequest
// 	if err := ctx.Bind(&objRequest); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err)
// 	}

// 	trvl := travelr.NewTraverPayouts("0088b3ed5e8e56d471d7f49ac4e2ee6f", "253183")
// 	rs, err := trvl.RoundTrip(objRequest.Segments)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err)
// 	}

// 	return ctx.JSON(http.StatusOK, rs)
// }

func main() {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	// gr := e.Group("/api/v1/public")
	// gr.POST("/round-trip", SearchRoundTrip)

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
