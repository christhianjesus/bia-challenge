package infrastructure

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	RegisterRoutes(*echo.Group)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SQLClient interface {
	QueryContext(ctx context.Context, query string, args ...any) (SQLRows, error)
}

type SQLRows interface {
	Close() error
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type DBAdapter struct {
	DB *sql.DB
}

func (adapter *DBAdapter) QueryContext(ctx context.Context, query string, args ...any) (SQLRows, error) {
	return adapter.DB.QueryContext(ctx, query, args...)
}
