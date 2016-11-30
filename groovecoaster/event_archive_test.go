package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestEventArchive(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		assetName string
		uri       string
		status    int
		success   bool
	}{
		{
			"../tests/assets/event_data.json",
			"mypage.groovecoaster.jp/sp/json/event_data.php?event_id=28&old_flag=true",
			200,
			true,
		},
		{
			"",
			"mypage.groovecoaster.jp/sp/json/event_data.php?event_id=28&old_flag=true",
			500,
			false,
		},
		{
			"../tests/assets/invalid_structure.json",
			"mypage.groovecoaster.jp/sp/json/event_data.php?event_id=28&old_flag=true",
			200,
			false,
		},
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
			scheme+test.uri,
			httpmock.NewBytesResponder(test.status, data),
		)

		_, err := testClient.EventArchive(28)
		if err != nil && test.success {
			t.Error(err)
		}
	}
}
