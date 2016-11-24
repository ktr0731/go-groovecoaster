package groovecoaster

import "fmt"

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
func (p *APIClient) MusicSummary() ([]*MusicSummary, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/music_list.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var ml musicSummary
	p.unmarshal(data, &ml)

	if ml.MusicSummary == nil {
		return nil, fmt.Errorf("invalid JSON structure")
	}

	return ml.MusicSummary, nil
}
