package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAPIClientMusicRankingPageCount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/music_ranking.json")
	if err != nil {
		t.Error(err.Error())
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/score_ranking_bymusic_bydifficulty.php?music_id=290&difficulty=0",
		httpmock.NewBytesResponder(200, data),
	)

	v, err := testClient.MusicRankingPageCount(290, Simple)
	if err != nil {
		t.Error(err)
	}
}
