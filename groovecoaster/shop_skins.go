package groovecoaster

// ShopSkins is the structure that represents shop sales info about skins
type ShopSkins struct {
	GC        int `json:",string"`
	ID        int
	Name      string
	New       StringToBool
	ProductID string `json:"product_id"`
}

type shopSkins struct {
	ShopSkins []*ShopSkins `json:"product_list"`
}

// ShopSkins fetches all skins in shop
func (p *APIClient) ShopSkins() ([]*ShopSkins, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/shop_skin_list.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var ss shopSkins
	p.unmarshal(data, &ss)

	return ss.ShopSkins, nil
}
