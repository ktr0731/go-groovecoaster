package groovecoaster

import (
	"fmt"
	"strconv"
)

// Rank is rank each difficulty
type Rank struct {
	Rank int
}

// Result is the structure that represent music result each difficulty
type Result struct {
	Rank      int
	IsClear   bool `json:"is_clear_mark"`
	IsFailed  bool `json:"is_failed_mark"`
	PlayCount int  `json:"play_count"`
	Adlib     int
	NoMiss    int `json:"no_miss"`
	FullChain int `json:"full_chain"`
	Perfect   int
	MaxChain  int `json:"max_chain"`
	Score     int
	Rating    string
}

// Detail .
type Detail struct {
	ID         string `json:"music_id"`
	Title      string `json:"music_title"`
	Artist     string
	HasEx      bool
	IsFavorite bool
	Skin       string `json:"skin_name"`
	Message    string
	Simple     *Result `json:"simple_result_data"`
	Normal     *Result `json:"normal_result_data"`
	Hard       *Result `json:"hard_result_data"`
	Extra      *Result `json:"extra_result_data"`
}

type tmp struct {
	MusicDetail struct {
		ExFlag int    `json:"ex_flag"`
		Rank   []Rank `json:"user_rank"`
	} `json:"music_detail"`
}

// MusicDetail .
type MusicDetail struct {
	Detail `json:"music_detail"`
}

// Music fetch a music detail by music id
func (c *APIClient) Music(id int) (*MusicDetail, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/music_detail.php?music_id="

	data, err := c.get(uri + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var md *MusicDetail
	var tmp tmp
	c.unmarshal(data, &md)
	c.unmarshal(data, &tmp)

	if md == nil {
		return nil, fmt.Errorf("Invalid JSON structure")
	}

	// TODO: Nilの場合がある
	md.Simple.Rank = tmp.MusicDetail.Rank[0].Rank
	md.Normal.Rank = tmp.MusicDetail.Rank[1].Rank
	md.Hard.Rank = tmp.MusicDetail.Rank[2].Rank

	if tmp.MusicDetail.ExFlag > 0 {
		md.HasEx = true
		md.Extra.Rank = tmp.MusicDetail.Rank[3].Rank
	} else {
		md.HasEx = false
	}

	return md, nil
}
