package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestShopItems(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/shop_item_list.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/shop_item_list.php",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.ShopItems()
	if err != nil {
		t.Error(err)
	}
}

func TestShopItems_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/shop_item_list.php",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.ShopItems()
	if err == nil {
		t.Error(err)
	}
}
