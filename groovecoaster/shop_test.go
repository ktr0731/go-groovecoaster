package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestShopSummary(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/shop_sales_data.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/shop_sales_data.php",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.ShopSummary()
	if err != nil {
		t.Error(err)
	}
}

func TestShopSummary_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/shop_sales_data.php",
		httpmock.NewStringResponder(200, `{"test": "test"}`),
	)

	_, err := testClient.ShopSummary()
	if err == nil {
		t.Error(err)
	}
}

func TestShopSummary_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/shop_sales_data.php",
		httpmock.NewStringResponder(500, `{"test": "test"}`),
	)

	_, err := testClient.ShopSummary()
	if err == nil {
		t.Error(err)
	}
}
