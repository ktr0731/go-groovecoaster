package groovecoaster

import (
	"fmt"
	"strconv"
)

// Difficulty represents difficulty each of musics
type Difficulty int

// Enum for difficulty
const (
	Simple Difficulty = iota
	Normal
	Hard
	Extra
)

// RankingElement is the structure that represents a ranking detail of a player
type RankingElement struct {
	Score            string `json:"event_point"`
	Name             string `json:"player_name"`
	LastPlayedArcade string `json:"last_play_tenpo_name"`
	Prefecture       string `json:"pref"`
	Rank             int
	Title            string
}

type ranking struct {
	Count   string
	Ranking []*RankingElement `json:"score_rank"`
}

// MusicRankingPageCount fetches last page number by music id and difficulty
func (c *APIClient) MusicRankingPageCount(id int, diff Difficulty) (int, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/score_ranking_bymusic_bydifficulty.php?music_id=%d&difficulty=%d"

	data, err := c.get(fmt.Sprintf(uri, id, diff))
	if err != nil {
		return -1, err
	}

	var rd *ranking
	c.unmarshal(data, &rd)

	if rd == nil {
		return -1, fmt.Errorf("Invalid JSON structure")
	}

	count, err := strconv.Atoi(rd.Count)
	if err != nil {
		return -1, fmt.Errorf("Invalid page count")
	}

	return count / 100, nil
}

// MusicRanking fetches a music score ranking by music id and difficulty
func (c *APIClient) MusicRanking(id int, diff Difficulty, page int) ([]*RankingElement, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/score_ranking_bymusic_bydifficulty.php?music_id=%d&difficulty=%d&page=%d"

	data, err := c.get(fmt.Sprintf(uri, id, diff, page))
	if err != nil {
		return nil, err
	}

	var rd *ranking
	c.unmarshal(data, &rd)

	if rd.Ranking == nil {
		return nil, fmt.Errorf("Invalid JSON structure")
	}
	fmt.Println(rd.Ranking[0])

	return rd.Ranking, nil
}
