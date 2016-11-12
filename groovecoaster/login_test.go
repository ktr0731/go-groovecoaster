package groovecoaster

import "testing"

func TestTryLogin(t *testing.T) {
	_, err := TryLogin()
	if err != nil {
		t.Error(err.Error())
	}
}
