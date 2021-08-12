package teams

import (
	"encoding/json"
	"io/ioutil"
)

// Team ...
type Team struct {
	ID     string `json:"ID"`
	Name   string `json:"name"`
	Location   string `json:"location"`
	Year   string `json:"year"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`	
	Ties   int    `json:"ties"`	
	Other  int    `json:"other"`  
	Games  int    `json:"games"`
	UUID   string `json:"uuid"`
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
