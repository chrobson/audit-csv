package main

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Audit represents pokerstars audyt structure
type Audit struct {
	DateTime          string `csv:"datetime"`
	Action            string `csv:"action"`
	Tournament        string `csv:"tournament"`
	Game              string `csv:"game"`
	Currency          string `csv:"currency"`
	Ammount           string `csv:"ammount"`
	Tmoney            string `csv:"tmoney"`
	Accuredstarscoins string `csv:"accuredstarscoins"`
	Wmoney            string `csv:"wmoney"`
	Balance           string `csv:"balance"`
	Totalstarcoins    string `csv:"totalstarcoins"`
	Tmoney1           string `csv:"tmoney1"`
	Wmoney1           string `csv:"wmoney1"`
}

// TODO: use gocsv

//ParseAudit process audit body into Audit struct
func ParseAudit(r io.Reader) ([]Audit, error) {
	var audits []Audit
	csvReader := csv.NewReader(r)
	csvReader.FieldsPerRecord = 13 // set other options
	if err := gocsv.UnmarshalCSV(csvReader, &audits); err != nil {
		return nil, err
	}

	return audits, nil
}

//saveAuditToDB saves audit into database
func saveAuditToDB(a []Audit) {
	db, err := sqlx.Open("postgres", "postgres://postgres:dbpass@localhost/TestDb?sslmode=disable")
	if err != nil {
		fmt.Println("Error in postgres connection: ", err)
	}

	for _, audit := range a {
		_, err = db.NamedExec(`INSERT INTO audit (datetime,action,tournament,game,currency,ammount,tmoney,accuredstarscoins,wmoney,balance,totalstarcoins,tmoney1,wmoney1) 
		VALUES (:datetime,:action,:tournament,:game,:currency,:ammount,:tmoney,:accuredstarscoins,:wmoney,:balance,:totalstarcoins,:tmoney1,:wmoney1)`, audit)
		if err != nil {
			fmt.Println("Error in postgres connection: ", err)
		}
	}
}

//countRB sums money and tournament money from audit
func countRb(playerInfos []Audit) (ammount, tmoney string) {
	// TODO: move this calculations to sql

	return "", ""
}
