package groovecoaster

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// TryLogin will attempt login to mypage in GrooveCoaster
// If login was successful, it return http.Client
func TryLogin() (*http.Client, error) {
	const loginURL = "https://mypage.groovecoaster.jp/sp/login/auth_con.php"

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	v := url.Values{}
	v.Add("nesicaCardId", os.Getenv("NESICA_CARD_ID"))
	v.Add("password", os.Getenv("NESICA_PASSWORD"))

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{Jar: jar}

	res, err := client.PostForm(loginURL, v)
	if err != nil {
		return nil, err
	}

	if url := res.Request.URL.String(); strings.Contains(url, "isError=true") {
		return nil, fmt.Errorf("Invalid NESiCA Card ID or password")
	} else if strings.Contains(url, "login_stop") {
		return nil, fmt.Errorf("Login attempts limit exceed")
	}

	return &client, nil
}
