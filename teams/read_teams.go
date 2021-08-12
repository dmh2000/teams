package teams

import (
	"encoding/json"
	"io/ioutil"
)

// Team ...
type Team struct {
	ID     string `json:"ID,omitempty"`
	Name   string `json:"name,omitempty"`
	Wins   int    `json:"wins,omitempty"`
	Losses int    `json:"losses,omitempty"`
	Ties   int    `json:"ties,omitempty"`
	Other  int    `json:"other,omitempty"`
	Games  int    `json:"games,omitempty"`
	UUID   string `json:"uuid,omitempty"`
}

// ReadTeam - reads the json file and returns a slice of data
func ReadTeams(fname string) ([]Team, error) {
	jsonBlob, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var tg []Team

	err = json.Unmarshal(jsonBlob, &tg)
	if err != nil {
		return nil, err
	}

	return tg, nil
}
