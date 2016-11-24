package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMusicList(t *testing.T) {
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

	_, err = testClient.MusicSummary()
	if err != nil {
		t.Error(err)
	}
}

func TestMusicList_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/music_list.php",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.MusicSummary()
	if err == nil {
		t.Error(err)
	}
}

func TestMusicList_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/music_list.php",
		httpmock.NewStringResponder(200, `{"test": "test"}`),
	)

	_, err := testClient.MusicSummary()
	if err == nil {
		t.Error(err)
	}
}
