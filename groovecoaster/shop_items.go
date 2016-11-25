package groovecoaster

// ShopItems is the structure that represents shop sales info about items
type ShopItems struct {
	ProductID  string `json:"product_id"`
	GC         int    `json:",string"`
	Possession int    `json:"possess,string"`
	ID         int
	Comment    string
	Name       string
	New        StringToBool
}

type shopItems struct {
	ShopItems []*ShopItems `json:"item_list"`
}

// ShopItems fetches all items in shop
func (c *APIClient) ShopItems() ([]*ShopItems, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/shop_item_list.php"

	data, err := c.get(uri)
	if err != nil {
		return nil, err
	}

	var si shopItems
	c.unmarshal(data, &si)

	return si.ShopItems, nil
}
