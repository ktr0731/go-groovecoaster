package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestShopSkins(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/shop_skin_list.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/shop_skin_list.php",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.ShopSkins()
	if err != nil {
		t.Error(err)
	}
}
