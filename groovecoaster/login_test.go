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

	httpmock.RegisterResponder(
		"POST",
		scheme+"mypage.groovecoaster.jp/sp/login/auth_con.php",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "")
			resp.Request = req
			url, _ := url.Parse("https://mypage.groovecoaster.jp/sp/#/")
			resp.Request.URL = url
			return resp, nil
		},
	)

	_, err := tryLogin()
	if err != nil {
		t.Error(err)
	}
}

func TestTryLogin_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"POST",
		scheme+"mypage.groovecoaster.jp/sp/login/auth_con.php",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "")
			resp.Request = req
			url, _ := url.Parse("https://mypage.groovecoaster.jp/sp/login/auth.php?isError=true&val=-1")
			resp.Request.URL = url
			return resp, nil
		},
	)

	_, err := tryLogin()
	if err == nil {
		t.Error(err)
	}
}

func TestTryLogin_LoginStop(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"POST",
		scheme+"mypage.groovecoaster.jp/sp/login/auth_con.php",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "")
			resp.Request = req
			url, _ := url.Parse("https://mypage.groovecoaster.jp/sp/login/auth.php?login_stop=true&val=-1")
			resp.Request.URL = url
			return resp, nil
		},
	)

	_, err := tryLogin()
	if err == nil {
		t.Error(err)
	}
}
