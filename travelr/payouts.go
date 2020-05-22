package travelr

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetSignatureFromRequest(e echo.Context) error {
	var obj DefaultRequest
	if err := e.Bind(&obj); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error)
	}

	rs, err := MakeSignatureToken(token, obj.Data)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error)
	}

	return e.JSON(http.StatusOK, map[string]string{
		"signature": rs,
	})
}
