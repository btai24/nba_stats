package api

type Scoreboard struct {
	Parameters Parameters  `json:"parameters"`
	ResultSets []ResultSet `json:"resultSets"`
}

type ResultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:rowSet`
}

type Parameters struct {
	GameDate  string `json:"GameDate"`
	LeagueID  string `json:"LeagueID"`
	DayOffset int    `json:"DayOffset"`
}
