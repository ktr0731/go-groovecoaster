package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMusic(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		assetName string
		status    int
		success   bool
	}{
		{"../tests/assets/music_detail.json", 200, true},
		{"../tests/assets/music_detail_not_played.json", 200, false},
		{"", 500, false},
		{"", 200, false},
	}

	for _, test := range tests {
		data := []byte("")
		if test.assetName != "" {
			var err error
			data, err = ioutil.ReadFile(test.assetName)
			if err != nil {
				t.Error(err)
			}
		}

		httpmock.RegisterResponder(
			"GET",
			scheme+"mypage.groovecoaster.jp/sp/json/music_detail.php?music_id=0",
			httpmock.NewBytesResponder(test.status, data),
		)

		_, err := testClient.Music(0)
		if err != nil && test.success {
			t.Error(err)
		}
	}

}
