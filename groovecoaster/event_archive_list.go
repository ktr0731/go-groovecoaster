package groovecoaster

import "fmt"

// EventArchive is the structure that represents a row of all events
type EventArchive struct {
	OpenDate  string `json:"open_date"`
	CloseDate string `json:"close_date"`
	EventID   int    `json:"event_id"`
	Title     string
}

type eventArchiveList struct {
	EventArchiveList []*EventArchive `json:"event_info_list"`
}

// EventArchiveList fetches all events that has been held until now
func (p *APIClient) EventArchiveList() ([]*EventArchive, error) {
	const uri = "mypage.groovecoaster.jp/sp/json/event_info_list.php"

	data, err := p.get(uri)
	if err != nil {
		return nil, err
	}

	var ea eventArchiveList
	p.unmarshal(data, &ea)

	if ea.EventArchiveList != nil {
		return nil, fmt.Errorf("invalid JSON structure: EventArchiveList()")
	}

	return ea.EventArchiveList, nil
}
