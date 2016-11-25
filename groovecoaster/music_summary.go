package groovecoaster

// MusicSummary is the structure that represents a row of music list
type MusicSummary struct {
	ID           int    `json:"music_id"`
	Title        string `json:"music_title"`
	PlayCount    int    `json:"play_count"`
	LastPlayTime string `json:"last_play_time"`
}

type musicSummary struct {
	MusicSummary []*MusicSummary `json:"music_list"`
}

// MusicSummary fetches all musics name by array
func (c *APIClient) MusicSummary() ([]*MusicSummary, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/music_list.php"

	data, err := c.get(uri)
	if err != nil {
		return nil, err
	}

	var ml musicSummary
	c.unmarshal(data, &ml)

	return ml.MusicSummary, nil
}
