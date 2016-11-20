package groovecoaster

import (
	"fmt"
	"strconv"
)

type intToBool bool

func (i *intToBool) UnmarshalJSON(bytes []byte) error {
	n, err := strconv.Atoi(string(bytes))
	if err != nil {
		return err
	}

	switch n {
	case 1:
		*i = true
	case 0:
		*i = false
	default:
		return fmt.Errorf("Invalid value in intToBool")
	}

	return nil
}

// Sales is the structure that represents shop sales info
type Sales struct {
	Avatar  intToBool `json:"avatar_sale"`
	Item    intToBool `json:"item_sale"`
	Message intToBool `json:"message_sale"`
	Music   intToBool `json:"music_sale"`
	Skin    intToBool `json:"skin_sale"`
	Sound   intToBool `json:"sound_sale"`
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
