package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestOnlineBattleSummary(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/online_battle_list.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/online_battle_list.php",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.OnlineBattleSummary()
	if err != nil {
		t.Error(err)
	}
}

func TestOnlineBattleSummary_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/online_battle_list.php",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.OnlineBattleSummary()
	if err == nil {
		t.Error(err)
	}
}
