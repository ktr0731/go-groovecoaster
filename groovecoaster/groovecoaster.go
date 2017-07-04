package groovecoaster

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// APIClient is the structure that represents Groove Coaster client
type APIClient struct {
	client *http.Client
}

// New creates a groovecoaster client
func New() *APIClient {
	client, err := tryLogin()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return &APIClient{
		client: client,
	}
}

func (c *APIClient) get(uri string) (io.ReadCloser, error) {
	res, err := c.client.Get("https://" + uri)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch %s failed (Status: %s)", uri, res.Status)
	}

	return res.Body, nil
}

func (c *APIClient) decode(w io.ReadCloser, i interface{}) error {
	if err := json.NewDecoder(w).Decode(i); err != nil {
		return fmt.Errorf("decode failed: %s", err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("writer closing failed: %s", err)
	}

	return nil
}
