package groovecoaster

// ShopSkins is the structure that represents shop sales info about skins
type ShopSkins struct {
	ProductID string `json:"product_id"`
	GC        int    `json:",string"`
	ID        int
	Name      string
	New       StringToBool
}

type shopSkins struct {
	ShopSkins []*ShopSkins `json:"product_list"`
}

// ShopSkins fetches all skins in shop
func (c *APIClient) ShopSkins() ([]*ShopSkins, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/shop_skin_list.php"

	data, err := c.get(uri)
	if err != nil {
		return nil, err
	}

	var ss shopSkins
	c.unmarshal(data, &ss)

	return ss.ShopSkins, nil
}
