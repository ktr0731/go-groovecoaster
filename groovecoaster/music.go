package groovecoaster

import "fmt"

// Result is the structure that represents music result each difficulty
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

// Detail is the structure that represents a music detail
type Detail struct {
	ID         string `json:"music_id"`
	Title      string `json:"music_title"`
	Artist     string
	ImageURL   string
	HasEx      IntToBool `json:"ex_flag"`
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
		Rank []struct {
			Rank int
		} `json:"user_rank"`
	} `json:"music_detail"`
}

// MusicDetail is a wrapper type of JSON
type MusicDetail struct {
	Detail `json:"music_detail"`
}

// Music fetches a music detail by music id
func (c *APIClient) Music(id int) (*MusicDetail, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/music_detail.php?music_id=%d"

	data, err := c.get(fmt.Sprintf(uri, id))
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

	if md.Simple != nil {
		md.Simple.Rank = tmp.MusicDetail.Rank[0].Rank
	}

	if md.Normal != nil {
		md.Normal.Rank = tmp.MusicDetail.Rank[1].Rank
	}

	if md.Hard != nil {
		md.Hard.Rank = tmp.MusicDetail.Rank[2].Rank
	}

	if md.HasEx && md.Extra != nil {
		md.Extra.Rank = tmp.MusicDetail.Rank[3].Rank
	}

	md.ImageURL = "https://mypage.groovecoaster.jp/sp/music/music_image.php?music_id=" + md.ID

	return md, nil
}
