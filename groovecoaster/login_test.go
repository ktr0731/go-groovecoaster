package groovecoaster

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestTryLogin(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		dist    string
		success bool
	}{
		{"https://mypage.groovecoaster.jp/sp/#/", true},
		{"https://mypage.groovecoaster.jp/sp/login/auth.php?isError=true&val=-1", false},
		{"https://mypage.groovecoaster.jp/sp/login/auth.php?isError=true&val=-1", false},
	}

	for _, test := range tests {
		httpmock.RegisterResponder(
			"POST",
			scheme+"mypage.groovecoaster.jp/sp/login/auth_con.php",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, "")
				resp.Request = req
				url, _ := url.Parse(test.dist)
				resp.Request.URL = url
				return resp, nil
			},
		)

		_, err := tryLogin()
		if err != nil && test.success {
			t.Error(err)
		}
	}
}
