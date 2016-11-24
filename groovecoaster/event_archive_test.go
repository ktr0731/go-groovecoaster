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
