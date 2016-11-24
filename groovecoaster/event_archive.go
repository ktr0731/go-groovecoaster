package groovecoaster

import "fmt"

// PlayerEventDetail is the structure that represents player's awards and scores
type PlayerEventDetail struct {
	Awards     *Awards `json:"award_data"`
	EventPoint int     `json:"event_point,string"`
	Score      int     `json:"high_score"`
	Rank       int
}

// EventDetail is the structure that represents the event result by event id
type EventDetail struct {
	Player    *PlayerEventDetail `json:"user_event_data"`
	OpenDate  string             `json:"open_date"`
	CloseDate string             `json:"close_date"`
	Title     string             `json:"title_name"`
}

type eventDetail struct {
	EventDetail *EventDetail `json:"event_data"`
}

// EventArchiveDetail fetches a event archive detail by event id
func (c *APIClient) EventArchiveDetail(eventID int) (*EventDetail, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/event_data.php?event_id=%d&old_flag=true"

	data, err := c.get(fmt.Sprintf(uri, eventID))
	if err != nil {
		return nil, err
	}

	var ed eventDetail
	c.unmarshal(data, &ed)

	if ed.EventDetail == nil {
		return nil, fmt.Errorf("Invalid JSON structure")
	}

	return ed.EventDetail, nil
}