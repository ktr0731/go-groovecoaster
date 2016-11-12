package groovecoaster

import (
	"fmt"
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
