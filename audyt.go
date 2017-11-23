package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DateTime struct {
	time.Time
}


// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006/01/02 3:04 PM", csv)
	if err != nil {
		return err
	}
	return nil
}

//Audit represents pokerstars audyt structure
type Audit struct {
	DateTime          DateTime `csv:"datetime"`
	Action            string   `csv:"action"`
	Tournament        string   `csv:"tournament"`
	Game              string   `csv:"game"`
	Currency          string   `csv:"currency"`
	Ammount           float64  `csv:"ammount"`
	Tmoney            float64  `csv:"tmoney"`
	Accuredstarscoins float64  `csv:"accuredstarscoins"`
	Wmoney            float64  `csv:"wmoney"`
	Balance           float64  `csv:"balance"`
	Totalstarcoins    float64  `csv:"totalstarcoins"`
	TTmoney           float64  `csv:"ttmoney"`
	WWmoney           float64  `csv:"wwmoney"`
}

//ParseAudit process audit body into Audit struct
func ParseAudit(r io.Reader) ([]Audit, error) {
	var audits []Audit
	csvReader := csv.NewReader(r)
	csvReader.LazyQuotes = true
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
		_, err = db.Exec(`INSERT INTO audit (datetime,action,tournament,game,currency,ammount,tmoney,accuredstarscoins,wmoney,balance,totalstarcoins,tmoney1,wmoney1) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`,
		audit.DateTime.Format("2006-01-02 15:04:05"),
		audit.Action,
		audit.Tournament,
		audit.Game,
		audit.Currency,
		audit.Ammount,
		audit.Tmoney,
		audit.Accuredstarscoins,
		audit.Wmoney,
		audit.Balance,
		audit.Totalstarcoins,
		audit.TTmoney,
		audit.WWmoney)
		if err != nil {
			fmt.Println("Error in postgres connection: ", err)
		}
	}
}

//countRB returns rakeback from specific dataframe
func countRb(from, to string) (ammount, tmoney float64) {
	db, err := sqlx.Open("postgres", "postgres://postgres:dbpass@localhost/TestDb?sslmode=disable")
	if err != nil {
		fmt.Println("Error in postgres connection: ", err)
	}
	rows, err := db.Query(fmt.Sprintf(`select sum(ammount) as ammount, sum(tmoney) as tmoney  from audit where action like '%s' and datetime between ('%s') and ('%s')`,"%Star%", from,to))

	for rows.Next() {
		if err := rows.Scan(&ammount, &tmoney); err != nil {
			fmt.Println("error ", err)
		}
	}
	return
}
