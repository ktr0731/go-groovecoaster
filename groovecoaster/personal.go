package groovecoaster

// Statistics is the structure that represent stage statistics of all musics
type Statistics struct {
	All       int
	Clear     int
	Fullchain int
	Nomiss    int
	Perfect   int
	S         int
	Ss        int
	Sss       int
}

// Personal is the structure that represent personal configration
type Personal struct {
	AverageScore      string `json:"average_score"`
	Avatar            string
	FriendApplication bool
	Level             int
	Name              string `json:"player_name"`
	Rank              int
	Title             string
	TotalMusic        int    `json:"total_music"`
	TotalPlayMusic    int    `json:"total_play_music"`
	TotalScore        string `json:"total_score"`
	TotalTrophy       string `json:"total_trophy"`
	TrophyRank        string `json:"trophy_rank"`
}

type playerData struct {
	Status     int
	Personal   *Personal   `json:"player_data"`
	Statistics *Statistics `json:"stage"`
}

// Personal fetch player profile
func (p *APIClient) Personal() (*Personal, error) {
	const personal = "mypage.groovecoaster.jp/sp/json/player_data.php"

	data, err := p.get(personal)
	if err != nil {
		return nil, err
	}

	var pd playerData
	p.unmarshal(data, &pd)

	return pd.Personal, nil
}

// Statistics fetch music statistics
func (p *APIClient) Statistics() (*Statistics, error) {
	const statistics = "mypage.groovecoaster.jp/sp/json/player_data.php"

	data, err := p.get(statistics)
	if err != nil {
		return nil, err
	}

	var pd playerData
	p.unmarshal(data, &pd)

	return pd.Statistics, nil
}
