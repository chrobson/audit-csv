package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
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

func processFile(csvFile multipart.File) (float32, float32) {
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
	var ammount, tmoney float32
	for _, playerInfo := range playerInfos {
		if strings.Contains(playerInfo.Action, "Stars") {
			ammount = ammount + playerInfo.Amount
			tmoney = tmoney + playerInfo.TMoney
		}

	}
	return ammount, tmoney
}
