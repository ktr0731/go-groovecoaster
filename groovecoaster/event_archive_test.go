package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestEventArchive(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/event_data.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_data.php?event_id=28&old_flag=true",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.EventArchiveDetail(28)
	if err != nil {
		t.Error(err)
	}
}

func TestEventArchive_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_data.php?event_id=28&old_flag=true",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.EventArchiveDetail(28)
	if err == nil {
		t.Error(err)
	}
}

func TestEventArchive_EmptyBody(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_data.php?event_id=28&old_flag=true",
		httpmock.NewStringResponder(200, ""),
	)

	_, err := testClient.EventArchiveDetail(28)
	if err == nil {
		t.Error(err)
	}
}

func TestEventArchive_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_data.php?event_id=28&old_flag=true",
		httpmock.NewStringResponder(200, `{"test": "test"}`),
	)

	_, err := testClient.EventArchiveDetail(28)
	if err == nil {
		t.Error(err)
	}
}
