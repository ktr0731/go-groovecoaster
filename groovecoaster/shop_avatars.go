package groovecoaster

// ShopAvatars is the structure that represents shop sales info about avatars
type ShopAvatars struct {
	GC        int `json:",string"`
	ID        int
	Name      string
	New       StringToBool
	ProductID string `json:"product_id"`
}

type shopAvatars struct {
	ShopAvatars []*ShopAvatars `json:"product_list"`
}

// ShopAvatars fetches all avatars in shop
func (p *APIClient) ShopAvatars() ([]*ShopAvatars, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/shop_avatar_list.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var sa shopAvatars
	p.unmarshal(data, &sa)

	return sa.ShopAvatars, nil
}
