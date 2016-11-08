package groovecoaster

import (
	"fmt"
	"io/ioutil"
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

	err := godotenv.Load()
	if err != nil {
		return nil, err
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
	fmt.Println("OK")

	res, err = client.Get("https://mypage.groovecoaster.jp/sp/json/player_data.php")
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	return &client, nil
}
