package main

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// AuditRequest ...
type AuditRequest struct {
	Audit string `json:"audit"`
}

type AuditHandler struct {
	//....
	db sql.DB
}

func NewAuditHandler(db sql.DB) *AuditHandler {
	return &AuditHandler{db}
}

func (h *AuditHandler) handle(c echo.Context) error {
	//....
	return nil
}

func auditHandler(c echo.Context) error {
	var req AuditRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	audits, err := ParseAudit(strings.NewReader(req.Audit))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	ammount, tmoney := countRb(audits)
	log.Info(ammount, tmoney)

	// store to database
	saveAuditToDB(audits)
	return nil
}

// func TestX(t testing.T) {
// 	// ......
// 	h := NewAuditHandler(db)
// 	if err := h.handle(c); err != nil {
// 		t.Fatalf("/v1/players/financials failed: %s", err)
// 	}
// }

type GetAuditResponse struct {
	Money  float64 `json:"money"`
	Tmoney float64 `json:"tmoney"`
}

func getAuditMoneyHandler(c echo.Context) error {
	// GET /audit/money?date=15-11-2017T00:00:00Z
	//u, err := c.Param("date")
	// if err != nil {
	// 	return err
	// }
	// u.Get("date")

	money, tmoney := 1.0, 1.0 // get from database

	var resp = GetAuditResponse{
		Money:  money,
		Tmoney: tmoney,
	}
	return c.JSON(http.StatusOK, resp)
}
