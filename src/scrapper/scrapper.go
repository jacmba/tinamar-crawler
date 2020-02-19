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
