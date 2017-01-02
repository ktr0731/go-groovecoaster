package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestShopSkins(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		assetName string
		status    int
		success   bool
	}{
		{"../tests/assets/shop_skin_list.json", 200, true},
		{"", 500, false},
	}

	for _, test := range tests {
		data := []byte("")

		if test.assetName != "" {
			var err error

			data, err = ioutil.ReadFile(test.assetName)
			if err != nil {
				t.Error(err)
			}
		}
		httpmock.RegisterResponder(
			"GET",
			scheme+"mypage.groovecoaster.jp/sp/json/shop_skin_list.php",
			httpmock.NewBytesResponder(test.status, data),
		)

		_, err := testClient.ShopSkins()
		if err != nil && test.success {
			t.Error(err)
		}
	}
}
