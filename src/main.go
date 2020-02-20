package main

import (
	"log"
	"strconv"
	"time"

	"./persistence"
	"./scrapper"
)

func exec(sc scrapper.Scrapper, pers persistence.Persistence) {
	period, _ := strconv.Atoi(executionPeriod)
	for {
		log.Println("---- STARTING CRAWLING PROCESS ----")
		body, getErr := sc.Get()

		if getErr != nil {
			panic(getErr)
		}

		league := sc.ExtractLeague(body)
		leagueHeader := sc.ExtractLeagueHeader(league)
		leagueFields := sc.ParseLeagueHeader(leagueHeader)
		leagueTeams := sc.ExtractLeagueTeams(league)
		leagueMap := sc.ParseLeagueTeams(leagueFields, leagueTeams)

		strBoardErr := pers.StoreLeaderBoard(leagueMap)

		if strBoardErr != nil {
			panic(strBoardErr)
		}
		log.Println("---- FINISHED CRAWLING PROCESS ----")

		time.Sleep(time.Duration(period) * time.Hour)
	}
}

func main() {
	log.Println("Starting crawling service")

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

	exec(sc, pers)
}
