package groovecoaster

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestShopMessages(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for i := 1; i <= 5; i++ {
		data, err := read(i)
		if err != nil {
			t.Error(err)
		}
		httpmock.RegisterResponder(
			"GET",
			fmt.Sprintf(scheme+"mypage.groovecoaster.jp/sp/#/sp_me/%d", i),
			httpmock.NewBytesResponder(200, data),
		)
	}

	_, err := testClient.ShopMessages()
	if err != nil {
		t.Error(err)
	}
}

func TestShopMessages_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for i := 1; i <= 5; i++ {
		httpmock.RegisterResponder(
			"GET",
			fmt.Sprintf(scheme+"mypage.groovecoaster.jp/sp/#/sp_me/%d", i),
			httpmock.NewStringResponder(500, ""),
		)
	}

	_, err := testClient.ShopMessages()
	if err == nil {
		t.Error(err)
	}
}

func read(i int) ([]byte, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("../tests/assets/shop_message_list_%d.json", i))
	if err != nil {
		return nil, err
	}

	return data, nil
}
