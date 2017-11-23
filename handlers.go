package main

import (
	"database/sql"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo"
	_"github.com/labstack/gommon/log"
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
	re := regexp.MustCompile(`(\d)(,)(\d{3})`)
	s := re.ReplaceAllString(req.Audit, "$1$3")
	r := strings.NewReader(s)
	audits, err := ParseAudit(r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

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
	dateFrom := c.QueryParam("start")
	dateTo := c.QueryParam("end")

	money, tmoney := countRb(dateFrom, dateTo)

	var resp = GetAuditResponse{
		Money:  money,
		Tmoney: tmoney,
	}
	return c.JSON(http.StatusOK, resp)
}
