package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestOnlineBattleDetail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/online_battle_detail.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/online_battle_detail.php?eid=34&mid=6448",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.OnlineBattleDetail("34", "6448")
	if err != nil {
		t.Error(err)
	}
}
