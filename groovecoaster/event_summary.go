package groovecoaster

import "fmt"

// EventResult is your play data in an event
type EventResult struct {
	EP        string `json:"event_point"`
	HighScore int    `json:"high_score"`
	Rank      int
}

// EventSummary is a summary of an event that is being held now
type EventSummary struct {
	*EventResult `json:"user_event_data"`
	StartDate    string `json:"open_date"`
	EndDate      string `json:"close_date"`
	Title        string `json:"title_name"`
}

type eventSummary struct {
	*EventSummary `json:"event_data"`
}

// EventSummary fetches a summary of event
func (c *APIClient) EventSummary() (*EventSummary, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/event_data.php"

	data, err := c.get(uri)
	if err != nil {
		return nil, err
	}

	var es eventSummary
	c.unmarshal(data, &es)

	if es.EventSummary == nil {
		return nil, fmt.Errorf("Invalid JSON structure")
	}

	return es.EventSummary, nil
}