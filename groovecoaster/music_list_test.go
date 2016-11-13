package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAPIMusicList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/music_list.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/music_list.php",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.MusicList()
	if err != nil {
		t.Error(err)
	}
}

func TestAPIMusicList_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/music_list.php",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.MusicList()
	if err == nil {
		t.Error(err)
	}
}
