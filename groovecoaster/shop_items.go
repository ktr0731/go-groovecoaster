package groovecoaster

// ShopItems is the structure that represents shop sales info about items
type ShopItems struct {
	Comment    string
	GC         int `json:",string"`
	ID         int
	Possession int `json:"possess,string"`
	Name       string
	New        StringToBool
	ProductID  string `json:"product_id"`
}

type shopItems struct {
	ShopItems []*ShopItems `json:"item_list"`
}

// ShopItems fetches all items in shop
func (p *APIClient) ShopItems() ([]*ShopItems, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/shop_item_list.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var si shopItems
	p.unmarshal(data, &si)

	return si.ShopItems, nil
}
