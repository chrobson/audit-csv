package main

import (
	"github.com/labstack/echo"
)

// NewServer creates servers, set routes, auth, logger and more.
func NewServer(cfg *Config) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.POST("/audit", auditHandler)
	e.GET("/audit/money", getAuditMoneyHandler)
	return e
}
