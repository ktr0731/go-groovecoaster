package groovecoaster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// APIClient is the structure that represents Groove Coaster client
type APIClient struct {
	client *http.Client
}

// New creates a groovecoaster client
func New() *APIClient {
	client, err := TryLogin()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	return &APIClient{
		client: client,
	}
}

func (c *APIClient) get(uri string) ([]byte, error) {
	res, err := c.client.Get("https://" + uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Fetch %s failed (Status: %d)", uri, res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("Response body is empty")
	}

	return data, nil
}

func (c *APIClient) unmarshal(data []byte, i interface{}) error {
	if err := json.Unmarshal(data, &i); err != nil {
		return fmt.Errorf("Unmarshal failed: %s", err)
	}

	return nil
}
