package nodesdkapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *NodeSDKHandler) Announce(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, nil)
}
