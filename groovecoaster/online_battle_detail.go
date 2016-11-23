package groovecoaster

import "fmt"

type playerDetail struct {
	Rank        int
	Title       string
	Avatar      string
	Difficulty1 StringToDifficulty `json:"difficulty_1st"`
	Difficulty2 StringToDifficulty `json:"difficulty_2nd"`
	Difficulty3 StringToDifficulty `json:"difficulty_3rd"`
	EntryNo     string             `json:"entry_no"`
	Item1       string             `json:"item_1st"`
	Item2       string             `json:"item_2nd"`
	Item3       string             `json:"item_3rd"`
	Music1      string             `json:"music_1st"`
	Music2      string             `json:"music_2nd"`
	Music3      string             `json:"music_3rd"`
	Name        string             `json:"player_name"`
	Prefecture  string             `json:"pref"`
	Star        int                `json:"result_star,string"`
	Arcade      string             `json:"tenpo_name"`
}

// OnlineBattleDetail is the structure that represents online battle result
type OnlineBattleDetail []*playerDetail

type onlineBattleDetail struct {
	OnlineBattleDetail OnlineBattleDetail
}

// OnlineBattleDetail fetches a online battle detail by eid and mid
func (c *APIClient) OnlineBattleDetail(eid string, mid string) (OnlineBattleDetail, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/online_battle_detail.php?eid=%s&mid=%s"

	data, err := c.get(fmt.Sprintf(uri, eid, mid))
	if err != nil {
		return nil, err
	}

	var obd onlineBattleDetail
	c.unmarshal(data, &obd)

	if obd.OnlineBattleDetail == nil {
		return nil, fmt.Errorf("Invalid JSON structure")
	}

	return obd.OnlineBattleDetail, nil
}
