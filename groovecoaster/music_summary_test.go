package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMusicList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		assetName string
		status    int
		success   bool
	}{
		{"../tests/assets/music_list.json", 200, true},
		{"", 500, false},
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
			scheme+"mypage.groovecoaster.jp/sp/json/music_list.php",
			httpmock.NewBytesResponder(test.status, data),
		)

		_, err := testClient.MusicSummary()
		if err != nil && test.success {
			t.Error(err)
		}
	}
}
