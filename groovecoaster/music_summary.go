package groovecoaster

// MusicSummary is the structure that represents a row of music list
type MusicSummary struct {
	ID           int    `json:"music_id"`
	Title        string `json:"music_title"`
	PlayCount    int    `json:"play_count"`
	LastPlayTime string `json:"last_play_time"`
}

// MusicSummary fetches all musics name by array
func (c *APIClient) MusicSummary() ([]MusicSummary, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/music_list.php"

	body, err := c.get(uri)
	if err != nil {
		return nil, err
	}

	var ms struct {
		MusicSummary []MusicSummary `json:"music_list"`
	}

	if err := c.decode(body, &ms); err != nil {
		return nil, err
	}

	return ms.MusicSummary, nil
}
