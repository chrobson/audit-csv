package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func testPage(c echo.Context) error {
	s := `<h1>Upload single file with fields</h1>
	<form action="/upload" method="post" enctype="multipart/form-data">
	Files: <input type="file" name="file"><br><br>
	<input type="submit" value="Submit">
	</form>`
	return c.HTML(http.StatusOK, fmt.Sprint(s))
}

func auditHandler(c echo.Context) error {

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	monney, tmoney := processFile(src)

	sHTML := `<p>File %s uploaded successfully!</p>
	Money from chests = %f <br/>
	TMoney from chests = %f <br/`
	return c.HTML(http.StatusOK, fmt.Sprintf(sHTML, file.Filename, monney, tmoney))
}
