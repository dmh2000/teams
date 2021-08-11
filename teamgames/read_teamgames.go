package teamgames

import (
	"encoding/json"
	"io/ioutil"
)

// TeamGames ...
type TeamGames struct {
	ID				string `json:"ID,omitempty"`
	Wins			int `json:"wins,omitempty"`
	Losses			int `json:"losses,omitempty"`
	Ties			int `json:"ties,omitempty"`
	Other			int `json:"other,omitempty"`
	Games			int `json:"games,omitempty"`
	UUID			string `json:"uuid,omitempty"`
}

// ReadTeamGames - reads the json file and returns a slice of data
func ReadTeamGames(fname string) ([]TeamGames, error) {
	jsonBlob, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var tg []TeamGames

	err = json.Unmarshal(jsonBlob, &tg)
	if err != nil {
		return nil, err
	}

	return tg, nil
}

