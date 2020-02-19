package main

import (
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

	fmt.Println(leagueFields)
	fmt.Println(leagueTeams)
}
