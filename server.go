package main

import (
	"github.com/labstack/echo"
)

type Audyt struct {
	DateTime          string  `json:"datetime"`
	Action            string  `json:"action"`
	Tournament        string  `json:"tournament"`
	Game              string  `json:"game"`
	Currency          string  `json:"currency"`
	Amount            float32 `json:"ammount"`
	AccruedStarsCoins string  `json:"accuredstarscoins"`
	TMoney            float32 `json:"tmoney"`
	WMoney            string  `json:"wmoney"`
	Balance           string  `json:"balance"`
	TotalStarCoins    string  `json:"totalstarcoins"`
	TMoney1           string  `json:"tmoney1"`
	WMoney1           string  `json:"wmoney1"`
}

// NewServer creates servers, set routes, auth, logger and more.
func NewServer(cfg *Config) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.GET("/", testPage)
	e.POST("/upload", auditHandler)
	return e
}
