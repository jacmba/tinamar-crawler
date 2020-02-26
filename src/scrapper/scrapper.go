/*
scrapper
*/
package scrapper

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Scrapper data structure to perform cawling operations
type Scrapper struct {
	URL string
}

/*
Get performs HTTP get request to Crawler URL
and returns body as string
*/
func (c *Scrapper) Get() (string, error) {
	resp, err := http.Get(c.URL)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		return "", err
	}

	return string(body), nil
}

/*
ExtractLeague returns league classification text from given body
*/
func (c *Scrapper) ExtractLeague(body string) string {
	beginMark := "<!-- begin classification -->"
	endMark := "<!-- end clasification-->"

	return extractSubtext(body, beginMark, endMark)
}

/*
ExtractLeagueHeader returns the league header text
*/
func (c *Scrapper) ExtractLeagueHeader(league string) string {
	beginMark := "<thead>"
	endMark := "</thead>"

	return extractSubtext(league, beginMark, endMark)
}

/*
ParseLeagueHeader returns a string list with header fields
*/
func (c *Scrapper) ParseLeagueHeader(header string) []string {
	fields := make([]string, 0)
	rows := extractRows(header)

	for _, row := range rows {
		b := "<th scope=\"col\">"
		e := "</th>"

		if strings.Contains(row, b) {
			field := extractSubtext(row, b, e)
			fields = append(fields, field)
		}
	}

	return fields[1:]
}

/*
ExtractLeagueTeams get teams info from leaderboard table
*/
func (c *Scrapper) ExtractLeagueTeams(league string) string {
	beginMark := "<tbody>"
	endMark := "</tbody>"

	return extractSubtext(league, beginMark, endMark)
}

/*
ParseLeagueTeams converts raw html table data into map
for further processing
*/
func (c *Scrapper) ParseLeagueTeams(header []string, teams string) []map[string]string {
	teamsList := strings.Split(teams, "</tr>")
	result := make([]map[string]string, 0)

	for _, rawTeam := range teamsList {
		idPreffix := "<tr class=\"position-relative\" data-team-id=\""
		idSuffix := "\">"
		idBegin := strings.Index(rawTeam, idPreffix)

		if idBegin >= 0 {
			team := make(map[string]string)
			team["id"] = extractSubtext(rawTeam, idPreffix, idSuffix)

			posPreffix := "<th class=\"text-center\" scope=\"row\">"
			posSuffix := "</th>"
			team["pos"] = extractSubtext(rawTeam, posPreffix, posSuffix)

			fieldsList := strings.Split(rawTeam, "</td>")
			fields := extractFields(fieldsList[1:])
			fillFields(team, header, fields)
			result = append(result, team)
		}
	}

	return result
}

/*
fillFields Fills up a map reference with gien fields and values
*/
func fillFields(obj map[string]string, keys []string, values []string) {
	for i, key := range keys {
		obj[key] = values[i]
	}
}

/*
extractSubtext returns text within defined marks
*/
func extractSubtext(text, beginMark, endMark string) string {
	begin := strings.Index(text, beginMark) + len(beginMark)
	end := strings.Index(text, endMark)

	return text[begin:end]
}

/*
extractFiels returns list of field values from given TD rows
*/
func extractFields(list []string) []string {
	result := make([]string, 0)

	for _, f := range list {
		preffix := "<td>"
		begin := strings.Index(f, preffix) + len(preffix)

		if begin >= 0 {
			field := f[begin:]
			result = append(result, field)
		}
	}

	return result
}

/*
extractRows returns a list of strings separating input string lines
*/
func extractRows(text string) []string {
	return strings.Split(text, "\n")
}

/*
ExtractLeagueFixtures returns the league fixtures info from html body
*/
func (c *Scrapper) ExtractLeagueFixtures(body string) string {
	fixturesPreffix := "<!-- begin content of calendar table 1-->"
	fixturesSuffix := "<!-- end calendar-->"
	return extractSubtext(body, fixturesPreffix, fixturesSuffix)
}

/*
SplitLeagueFixtures returns array with separated fixtures html text
*/
func (c *Scrapper) SplitLeagueFixtures(text string) []string {
	fixtureSeparator := "<div class=\"table-header w-100\">"
	return strings.Split(text, fixtureSeparator)[1:]
}

/*
ExtractFixture returns fixture number and raw games text string
*/
func (c *Scrapper) ExtractFixture(f string) (string, []string) {
	fixturePreffix := "<span class=\"d-inline-block\">Jornada "
	fixtureSuffix := "</span>"
	gamesPreffix := "<tbody>"
	gamesSuffix := "</tbody>"

	fixture := extractSubtext(f, fixturePreffix, fixtureSuffix)
	games := extractSubtext(f, gamesPreffix, gamesSuffix)
	gamesList := strings.Split(games, "<tr>")[1:]

	return fixture, gamesList
}

/*
ParseGame returns key-value map with game info from raw html
*/
func (c *Scrapper) ParseGame(f string, g string) map[string]string {
	game := make(map[string]string)
	game["fixture"] = f

	dateBegin := "data-date=\""
	timeBegin := "\" data-time=\""
	venueBegin := "\" data-place=\""
	dataEnd := "\"><i class=\"fa fa-calendar\""

	game["date"] = extractSubtext(g, dateBegin, timeBegin)
	game["time"] = extractSubtext(g, timeBegin, venueBegin)
	game["venue"] = extractSubtext(g, venueBegin, dataEnd)

	rawData := strings.Split(g, "\n")
	rawHomeTeam := rawData[5]
	rawHomeScore := rawData[6]
	rawAwayScore := rawData[7]
	rawAwayTeam := rawData[8]
	rawDataBegin := "\">"
	rawDataEnd := "</td>"

	game["homeTeam"] = extractSubtext(rawHomeTeam, rawDataBegin, rawDataEnd)
	game["homeScore"] = extractSubtext(rawHomeScore, rawDataBegin, rawDataEnd)
	game["awayTeam"] = extractSubtext(rawAwayTeam, rawDataBegin, rawDataEnd)
	game["awayScore"] = extractSubtext(rawAwayScore, rawDataBegin, rawDataEnd)

	return game
}

/*
ParseGames converts array of raw html text to array of key-value game maps
*/
func (c *Scrapper) ParseGames(f string, gs []string) []map[string]string {
	games := make([]map[string]string, 0)

	for _, g := range gs {
		game := c.ParseGame(f, g)
		games = append(games, game)
	}

	return games
}

/*
ParseFixtures returns parser map of games from raw fixture html list
*/
func (c *Scrapper) ParseFixtures(fs []string) []map[string]string {
	games := make([]map[string]string, 0)

	for _, f := range fs {
		fixture, rawGames := c.ExtractFixture(f)
		parsedGames := c.ParseGames(fixture, rawGames)
		games = append(games, parsedGames...)
	}

	return games
}
