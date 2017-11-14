package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func auditHandler(c echo.Context) error {

	//-----------
	// Read file
	//-----------

	body, _ := ioutil.ReadAll(c.Request().Body)
	r := bytes.NewReader(body)
	monney, tmoney := processFile(r)

	sHTML := `<p>File uploaded successfully!</p>
	Money from chests = %f <br/>
	TMoney from chests = %f <br/`
	return c.HTML(http.StatusOK, fmt.Sprintf(sHTML, monney, tmoney))
}
