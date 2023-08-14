package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	RegisterRoutes(*echo.Group)
}
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
