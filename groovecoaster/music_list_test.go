package groovecoaster

import (
	"fmt"
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

	p, err := testClient.MusicList()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(p)
}
