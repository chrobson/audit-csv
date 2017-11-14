package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//Audyt represents pokerstars audyt structure
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

//Process audit file into JsonData
func processFileToJSON(csvFile io.Reader) []byte {
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var playerInfos []Audyt
	//Skip first three rows
	reader.Read()
	reader.Read()
	reader.Read()
	reader.FieldsPerRecord = 13
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		ammount, _ := strconv.ParseFloat(line[5], 32)
		tmoney, _ := strconv.ParseFloat(line[7], 32)
		playerInfos = append(playerInfos, Audyt{
			DateTime:          line[0],
			Action:            line[1],
			Tournament:        line[2],
			Game:              line[3],
			Currency:          line[4],
			Amount:            float32(ammount),
			AccruedStarsCoins: line[6],
			TMoney:            float32(tmoney),
			WMoney:            line[8],
			Balance:           line[9],
			TotalStarCoins:    line[10],
			TMoney1:           line[11],
			WMoney1:           line[12],
		})
	}
	jsondata, err := json.Marshal(playerInfos)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return jsondata
}

//countRB sums money and tournament money from audit
func countRb(jsonData []byte) (ammount, tmoney float32) {
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
