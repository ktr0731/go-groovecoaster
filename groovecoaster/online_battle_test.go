package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestOnlineBattle(t *testing.T) {
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

	_, err = testClient.OnlineBattle(34, 6448)
	if err != nil {
		t.Error(err)
	}
}

func TestOnlineBattle_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/online_battle_detail.php?eid=34&mid=6448",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.OnlineBattle(34, 6448)
	if err == nil {
		t.Error(err)
	}
}
