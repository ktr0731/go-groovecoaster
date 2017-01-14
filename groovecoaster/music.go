package groovecoaster

import "fmt"

// Result is the structure that represents music result each difficulty
type Result struct {
	IsClear   bool `json:"is_clear_mark"`
	IsFailed  bool `json:"is_failed_mark"`
	PlayCount int  `json:"play_count"`
	NoMiss    int  `json:"no_miss"`
	FullChain int  `json:"full_chain"`
	MaxChain  int  `json:"max_chain"`
	Rank      int
	Adlib     int
	Perfect   int
	Score     int
	Rating    string
}

// MusicDetail is the structure that represents a music detail
type MusicDetail struct {
	Simple     *Result   `json:"simple_result_data"`
	Normal     *Result   `json:"normal_result_data"`
	Hard       *Result   `json:"hard_result_data"`
	Extra      *Result   `json:"extra_result_data"`
	HasEx      IntToBool `json:"ex_flag"`
	ID         string    `json:"music_id"`
	Title      string    `json:"music_title"`
	Skin       string    `json:"skin_name"`
	Artist     string
	ImageURL   string
	IsFavorite bool
	Message    string
}

type tmp struct {
	MusicDetail struct {
		Rank []struct {
			Rank int
		} `json:"user_rank"`
	} `json:"music_detail"`
}

type musicDetail struct {
	MusicDetail MusicDetail `json:"music_detail"`
}

// Music fetches a music detail by music id
func (c *APIClient) Music(id int) (*MusicDetail, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/music_detail.php?music_id=%d"

	data, err := c.get(fmt.Sprintf(uri, id))
	if err != nil {
		return nil, err
	}

	var mdd musicDetail
	var tmp tmp
	c.unmarshal(data, &mdd)
	c.unmarshal(data, &tmp)

	md := mdd.MusicDetail

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

	return &md, nil
}
