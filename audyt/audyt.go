package audyt

import (
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

//Audyt represents pokerstars audyt structure
type Audyt struct {
	DateTime          time.Time `csv:"datetime"`
	Action            string    `csv:"action"`
	Tournament        string    `csv:"tournament"`
	Game              string    `csv:"game"`
	Currency          string    `csv:"currency"`
	Amount            float64   `csv:"ammount"`
	TMoney            float64   `csv:"tmoney"`
	AccruedStarsCoins float64   `csv:"accuredstarscoins"`
	WMoney            float64   `csv:"wmoney"`
	Balance           float64   `csv:"balance"`
	TotalStarCoins    float64   `csv:"totalstarcoins"`
	TMoney1           float64   `csv:"tmoney1"`
	WMoney1           float64   `csv:"wmoney1"`
}

// TODO: use gocsv

// Process audit file into JsonData
func ParseAudit(r io.Reader) ([]*Audyt, error) {
	var audits []*Audyt
	if err := gocsv.Unmarshal(r, &audits); err != nil {
		return nil, err
	}
	return audits, nil

	// reader := csv.NewReader(bufio.NewReader(r))
	// var playerInfos []Audyt
	// //Skip first three rows
	// reader.Read()
	// reader.Read()
	// reader.Read()
	// reader.FieldsPerRecord = 13
	// for {
	// 	line, error := reader.Read()
	// 	if error == io.EOF {
	// 		break
	// 	} else if error != nil {
	// 		log.Fatal(error)
	// 	}

	// 	ammount, _ := strconv.ParseFloat(line[5], 32)
	// 	tmoney, _ := strconv.ParseFloat(line[7], 32)
	// 	playerInfos = append(playerInfos, Audyt{
	// 		DateTime:          line[0],
	// 		Action:            line[1],
	// 		Tournament:        line[2],
	// 		Game:              line[3],
	// 		Currency:          line[4],
	// 		Amount:            float32(ammount),
	// 		AccruedStarsCoins: line[6],
	// 		TMoney:            float32(tmoney),
	// 		WMoney:            line[8],
	// 		Balance:           line[9],
	// 		TotalStarCoins:    line[10],
	// 		TMoney1:           line[11],
	// 		WMoney1:           line[12],
	// 	})
	// }
	// jsondata, err := json.Marshal(playerInfos)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// return jsondata
}

//countRB sums money and tournament money from audit
func countRb(jsonData []byte) (ammount, tmoney float32) {
	// TODO: move this calculations to sql
	var playerInfos []Audyt
	json.Unmarshal(jsonData, &playerInfos)
	for _, playerInfo := range playerInfos {
		if strings.Contains(playerInfo.Action, "Stars") {
			ammount = ammount + playerInfo.Amount
			tmoney = tmoney + playerInfo.TMoney
		}
	}
	return ammount, tmoney
}
