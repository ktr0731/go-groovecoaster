package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAPIClientMusic(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/music_detail.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/music_detail.php?music_id=0",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.Music(0)
	if err != nil {
		t.Error(err)
	}
}
