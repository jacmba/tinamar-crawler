package model

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

/*
Team - class to hold team info within the league table
*/
type Team struct {
	id      int
	name    string
	pos     int
	points  int
	played  int
	won     int
	draw    int
	lost    int
	favour  int
	against int
}

/*
GetID - Getter function to return team ID
*/
func (t *Team) GetID() int {
	return t.id
}

/*
ToBson Converts a team into BSON representation
*/
func (t *Team) ToBson() bson.M {
	return bson.M{
		"id":      t.id,
		"name":    t.name,
		"pos":     t.pos,
		"points":  t.points,
		"played":  t.played,
		"won":     t.won,
		"draw":    t.draw,
		"lost":    t.lost,
		"favour":  t.favour,
		"against": t.against,
	}
}

/*
MakeTeam - Method to construct a Team object from provided map
*/
func MakeTeam(m map[string]string) Team {
	atoi := strconv.Atoi
	id, _ := atoi(m["id"])
	name := m["EQUIPO"]
	pos, _ := atoi(m["pos"])
	points, _ := atoi(m["PTS"])
	played, _ := atoi(m["PJ"])
	won, _ := atoi(m["PG"])
	draw, _ := atoi(m["PE"])
	lost, _ := atoi(m["PP"])
	favour, _ := atoi(m["GF"])
	against, _ := atoi(m["GC"])

	return Team{
		id,
		name,
		pos,
		points,
		played,
		won,
		draw,
		lost,
		favour,
		against,
	}
}

/*
MapTeams - Convert array of map object into Team instances
*/
func MapTeams(ms []map[string]string) []Team {
	result := make([]Team, 0)

	for _, entry := range ms {
		team := MakeTeam(entry)
		result = append(result, team)
	}

	return result
}
