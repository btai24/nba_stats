package api

// type Scoreboard struct {
// 	Parameters Parameters  `json:"parameters"`
// 	ResultSets []ResultSet `json:"resultSets"`
// }
//
// type ResultSet struct {
// 	Name    string          `json:"name"`
// 	Headers []string        `json:"headers"`
// 	RowSet  [][]interface{} `json:rowSet`
// }
//
// type Parameters struct {
// 	GameDate  string `json:"GameDate"`
// 	LeagueID  string `json:"LeagueID"`
// 	DayOffset int    `json:"DayOffset"`
// }

type Scoreboard struct {
	NumGames int    `json:"numGames"`
	Games    []Game `json:"games"`
}

type Game struct {
	Id               string   `json:"gameId"`
	Arena            Arena    `json:"arena"`
	StartTimeEastern string   `json:"startTimeEastern"`
	Clock            string   `json:"clock"`
	Attendance       string   `json:"attendance"`
	Duration         Duration `json:gameDuration`
	Period           Period   `json:"Period"`
	VisitingTeam     Team     `json:"vTeam"`
	HomeTeam         Team     `json:"hTeam"`
}

type Arena struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type Duration struct {
	Hours   string `json:"hours"`
	Minutes string `json:"minutes"`
}

type Period struct {
	Current       int  `json:"current"`
	IsHalftime    bool `json:"isHalftime"`
	IsEndOfPeriod bool `json:"isEndOfPeriod"`
}

type Team struct {
	Id           string  `json:"teamId"`
	Abbreviation string  `json:"triCode"`
	Wins         string  `json:"win"`
	Losses       string  `json:"loss"`
	Score        string  `json:"score"`
	LineScore    []Score `json:"linescore"`
}

type Score struct {
	Score string `json:"score"`
}
