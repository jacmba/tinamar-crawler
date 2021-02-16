package model

import (
	"strconv"
	"strings"
)

/*
Fixture - data structure to hold games info
*/
type Fixture struct {
	Round     int    `bson:"round"`
	Venue     string `bson:"venue"`
	DateTime  string `bson:"datetime"`
	HomeTeam  string `bson:"home_team"`
	AwayTeam  string `bson:"away_team"`
	HomeScore int    `bson:"home_score"`
	AwayScore int    `bson:"away_score"`
	Played    bool   `bson:"played"`
}

/*
MakeFixture creates Fixture instance from key-value map
*/
func MakeFixture(f map[string]string) Fixture {
	atoi := strconv.Atoi
	round, _ := atoi(f["fixture"])
	venue, _ := f["venue"]
	homeTeam := f["homeTeam"]
	awayTeam := f["awayTeam"]
	homeScore, _ := atoi(f["homeScore"])
	awayScore, _ := atoi(f["awayScore"])
	played := f["homeScore"] != "-" && f["awayScore"] != "-"

	rawDate := strings.Split(f["date"], "/")
	d := rawDate[0]
	m := rawDate[1]
	y := rawDate[2]
	date := y + m + d

	rawTime := strings.Split(f["time"], ":")
	h := rawTime[0]
	mm := rawTime[1]
	time := h + mm

	datetime := date + "T" + time + "00"

	return Fixture{
		Round:     round,
		Venue:     venue,
		DateTime:  datetime,
		HomeTeam:  homeTeam,
		AwayTeam:  awayTeam,
		HomeScore: homeScore,
		AwayScore: awayScore,
		Played:    played,
	}
}

/*
MapFixtures returns Fixture business object list from array of parseable maps
*/
func MapFixtures(fs []map[string]string) []Fixture {
	fixtures := make([]Fixture, 0)

	for _, f := range fs {
		fixture := MakeFixture(f)
		fixtures = append(fixtures, fixture)
	}

	return fixtures
}
