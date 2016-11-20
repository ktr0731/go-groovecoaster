package groovecoaster

import "fmt"

// Sales is the structure that represents shop sales info
type Sales struct {
	Avatar  IntToBool `json:"avatar_sale"`
	Item    IntToBool `json:"item_sale"`
	Message IntToBool `json:"message_sale"`
	Music   IntToBool `json:"music_sale"`
	Skin    IntToBool `json:"skin_sale"`
	Sound   IntToBool `json:"sound_sale"`
}

// Shop is the structure that contains sales info, whether is openning
type Shop struct {
	Sales  *Sales `json:"shop_data"`
	IsOpen bool   `json:"open_flg"`
	Coin   int    `json:"current_coin"`
}

// Shop fetches shop information
func (c *APIClient) Shop() (*Shop, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/shop_sales_data.php"

	data, err := c.get(uri)
	if err != nil {
		return nil, err
	}

	var s *Shop
	c.unmarshal(data, &s)

	if s.Sales == nil {
		return nil, fmt.Errorf("Invalid JSON structure")
	}

	return s, nil
}
