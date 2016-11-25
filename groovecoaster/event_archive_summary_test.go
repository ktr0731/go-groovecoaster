package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestEventArchiveSummary(t *testing.T) {
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

	_, err = testClient.EventArchiveSummary()
	if err != nil {
		t.Error(err)
	}
}

func TestEventArchiveSummary_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_info_list.php",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.EventArchiveSummary()
	if err == nil {
		t.Error(err)
	}
}

func TestEventArchiveSummary_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_info_list.php",
		httpmock.NewStringResponder(200, `{"test": "test"}`),
	)

	_, err := testClient.EventArchiveSummary()
	if err == nil {
		t.Error(err)
	}
}
