package main

import (
	"encoding/json"
	"fmt"

	"./scrapper"
)

const tinamarURL = "http://ligatinamar.com/category/once_veteranos_38b"

func main() {
	fmt.Println("Starting crawling service")

	sc := scrapper.Scrapper{
		URL: tinamarURL,
	}

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
