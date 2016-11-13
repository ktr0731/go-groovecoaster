package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAPIClientPersonal(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/player_data.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/player_data.php",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.Personal()
	if err != nil {
		t.Error(err)
	}
}

func TestAPIClientPersonal_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/player_data.php",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.Personal()
	if err == nil {
		t.Error(err)
	}
}

func TestAPIClientPersonal_EmptyBody(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/player_data.php",
		httpmock.NewStringResponder(200, ""),
	)

	_, err := testClient.Personal()
	if err == nil {
		t.Error(err)
	}
}

func TestAPIClientPersonal_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/player_data.php",
		httpmock.NewStringResponder(200, `{"test": "test"}`),
	)

	_, err := testClient.Personal()
	if err == nil {
		t.Error(err)
	}
}

func TestAPIClientStatistics(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/player_data.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/player_data.php",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.Statistics()
	if err != nil {
		t.Error(err)
	}
}

func TestAPIClientStatistics_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/player_data.php",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.Statistics()
	if err == nil {
		t.Error(err)
	}
}

func TestAPIClientStatistics_EmptyBody(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/player_data.php",
		httpmock.NewStringResponder(200, ""),
	)

	_, err := testClient.Statistics()
	if err == nil {
		t.Error(err)
	}
}

func TestAPIClientStatistics_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/player_data.php",
		httpmock.NewStringResponder(200, `{"test": "test"}`),
	)

	_, err := testClient.Statistics()
	if err == nil {
		t.Error(err)
	}
}
