package groovecoaster

import "fmt"

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
	Personal   *Personal   `json:"player_data"`
	Statistics *Statistics `json:"stage"`
}

// Personal fetch player profile
func (p *APIClient) Personal() (*Personal, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/player_data.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var pd playerData
	p.unmarshal(data, &pd)

	if pd.Personal == nil {
		return nil, fmt.Errorf("Invalid JSON structure")
	}

	return pd.Personal, nil
}

// Statistics fetch music statistics
func (p *APIClient) Statistics() (*Statistics, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/player_data.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var pd playerData
	p.unmarshal(data, &pd)

	if pd.Statistics == nil {
		return nil, fmt.Errorf("Invalid JSON structure")
	}

	return pd.Statistics, nil
}
