package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestEventArchiveList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/event_info_list.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_info_list.php",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.EventArchiveList()
	if err != nil {
		t.Error(err)
	}
}

func TestEventArchiveList_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_info_list.php",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.EventArchiveList()
	if err == nil {
		t.Error(err)
	}
}

func TestEventArchiveList_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_info_list.php",
		httpmock.NewStringResponder(200, `{"test": "test"}`),
	)

	_, err := testClient.EventArchiveList()
	if err == nil {
		t.Error(err)
	}
}
