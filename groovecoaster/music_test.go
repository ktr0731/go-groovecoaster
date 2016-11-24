package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMusic(t *testing.T) {
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

func TestMusic_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/music_detail.php?music_id=0",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.Music(0)
	if err == nil {
		t.Error(err)
	}
}

func TestMusic_NotPlayed(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/music_detail_not_played.json")
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
