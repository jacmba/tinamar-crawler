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

	return fields
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
extractSubtext returns text within defined marks
*/
func extractSubtext(text, beginMark, endMark string) string {
	begin := strings.Index(text, beginMark) + len(beginMark)
	end := strings.Index(text, endMark)

	return text[begin:end]
}

/*
extractRows returns a list of strings separating input string lines
*/
func extractRows(text string) []string {
	return strings.Split(text, "\n")
}
