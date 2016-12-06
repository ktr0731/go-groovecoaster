package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

type testSuit struct {
	assetName string
	status    int
	success   bool
}

func TestMusicRankingPageCount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []testSuit{
		{"../tests/assets/music_ranking.json", 200, true},
		{"", 500, false},
		{"", 200, false},
	}

	for _, test := range tests {
		data := []byte("")

		if test.assetName != "" {
			var err error
			data, err = ioutil.ReadFile(test.assetName)
			if err != nil {
				t.Error(err.Error())
			}
		}

		httpmock.RegisterResponder(
			"GET",
			scheme+"mypage.groovecoaster.jp/sp/json/score_ranking_bymusic_bydifficulty.php?music_id=290&difficulty=0",
			httpmock.NewBytesResponder(test.status, data),
		)

		_, err := testClient.MusicRankingPageCount(290, Simple)
		if err != nil && test.success {
			t.Error(err)
		}

	}
}

func TestMusicRanking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []testSuit{
		{"../tests/assets/music_ranking.json", 200, true},
		{"", 500, false},
		{"", 200, false},
	}

	for _, test := range tests {
		data := []byte("")

		if test.assetName != "" {
			var err error
			data, err = ioutil.ReadFile(test.assetName)
			if err != nil {
				t.Error(err.Error())
			}
		}

		httpmock.RegisterResponder(
			"GET",
			scheme+"mypage.groovecoaster.jp/sp/json/score_ranking_bymusic_bydifficulty.php?music_id=290&difficulty=0&page=0",
			httpmock.NewBytesResponder(test.status, data),
		)

		_, err := testClient.MusicRanking(290, Simple, 0)
		if err != nil && test.success {
			t.Error(err)
		}
	}

}
