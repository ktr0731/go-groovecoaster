package groovecoaster

// Player is the structure that represents a part of players in online battle result
type Player struct {
	Name       string `json:"player_name"`
	Prefecture string `json:"pref"`
}

// OnlineBattleSummary is the structure that represents a row of online battle list
type OnlineBattleSummary struct {
	Date       string
	EID        string
	IsFavorite StringToBool `json:"fav"`
	MID        string
	Rank       int
	Star       int      `json:",string"`
	Players    []Player `json:"detail"`
}

type onlineBattleSummary struct {
	OnlineBattleSummary []*OnlineBattleSummary `json:"onlineBattleList"`
}

// OnlineBattleSummary fetches all online battle results
func (p *APIClient) OnlineBattleSummary() ([]*OnlineBattleSummary, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/online_battle_list.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var obs onlineBattleSummary
	p.unmarshal(data, &obs)

	return obs.OnlineBattleSummary, nil
}
