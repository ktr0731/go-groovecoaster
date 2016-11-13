package groovecoaster

import "fmt"

// Music .
type Music struct {
	MusicID      int    `json:"music_id"`
	MusicTitle   string `json:"music_title"`
	PlayCount    int    `json:"play_count"`
	LastPlayTime string `json:"last_play_time"`
}

type musicList struct {
	MusicList []Music `json:"music_list"`
}

// MusicList fetch all music name by array
func (p *APIClient) MusicList() ([]Music, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/music_list.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var ml musicList
	p.unmarshal(data, &ml)

	if len(ml.MusicList) == 0 {
		return nil, fmt.Errorf("Music not found")
	}

	return ml.MusicList, nil
}
