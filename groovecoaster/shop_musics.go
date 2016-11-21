package groovecoaster

import "fmt"

// ShopMusics is the structure that represents shop sales info about musics
type ShopMusics struct {
	Artist    string
	GC        StringToInt
	ID        string
	Name      string
	New       StringToBool
	ProductID string `json:"product_id"`
}

type shopMusics struct {
	ShopMusics []*ShopMusics `json:"product_list"`
	Status     int
}

// ShopMusics fetches all musics in shop
func (p *APIClient) ShopMusics() ([]*ShopMusics, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/shop_music_list.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var sm shopMusics
	p.unmarshal(data, &sm)

	fmt.Printf("%#v\n", sm.ShopMusics[0])
	return sm.ShopMusics, nil
}
