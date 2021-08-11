package teamgames

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Team struct {
	Abbr      string `json:"Abbr,omitempty"`
	League    string `json:"League,omitempty"`
	City      string `json:"City,omitempty"`
	Nickname  string `json:"Nickname,omitempty"`
	FirstYear string `json:"FirstYear,omitempty"`
	LastYear  string `json:"LastYear,omitempty"`
	UUID string `json:"UUID,omitempty"`
}

// ReadTeams - read the json fie and returns a slice of data
func ReadTeams(fname string) ([]Team, error) {
	jsonBlob, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var tm []Team

	err = json.Unmarshal(jsonBlob, &tm)
	if err != nil {
		return nil, err
	}

	return tm, nil
}

func nameFromAbbr(teams []Team, abbr string) string {
	for _,t:=range(teams) {
		if abbr == t.Abbr {
			return t.Nickname
		}
	}
	return "not found"
}

func MergeTeams(teams string, teamgames string) error {
	tg,err := ReadTeamGames(teamgames);
	if err != nil {
		return err
	}
	tm, err := ReadTeams(teams);
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i,_ := range(tg) {
		tg[i].Name = nameFromAbbr(tm,tg[i].ID)
	}

	b,err := json.Marshal(tg)
	if err != nil {
		return err;
	}

	fmt.Print(string(b))
	fmt.Println();

	return nil
}