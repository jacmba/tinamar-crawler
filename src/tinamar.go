package main

import (
	"encoding/json"
	"fmt"
	"log"

	"./persistence"
	"./scrapper"
)

const tinamarURL = "http://ligatinamar.com/category/once_veteranos_38b"
const mongoURL = "mongodb://localhost:27017"

func main() {
	fmt.Println("Starting crawling service")

	pers := persistence.Persistence{
		URI: mongoURL,
	}

	sc := scrapper.Scrapper{
		URL: tinamarURL,
	}

	dbErr := pers.Connect()

	if dbErr != nil {
		panic(dbErr)
	}

	log.Println("Connected to DB")

	body, getErr := sc.Get()

	if getErr != nil {
		panic(getErr)
	}

	league := sc.ExtractLeague(body)
	leagueHeader := sc.ExtractLeagueHeader(league)
	leagueFields := sc.ParseLeagueHeader(leagueHeader)
	leagueTeams := sc.ExtractLeagueTeams(league)
	leagueMap := sc.ParseLeagueTeams(leagueFields, leagueTeams)

	leagueJSON, _ := json.Marshal(leagueMap)
	fmt.Println(string(leagueJSON))
}
