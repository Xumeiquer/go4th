package go4th

import (
	"encoding/json"
	"fmt"
	"time"
)

type Alert struct {
	ID           string      `json:"_id,omitempty"`
	Typ          string      `json:"_type,omitempty"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Severity     Severity    `json:"severity"`
	Date         time.Time   `json:"date"`
	Tags         []string    `json:"tags"`
	TLP          TLP         `json:"tlp"`
	Status       AlertStatus `json:"status"`
	Type         string      `json:"type"`
	Source       string      `json:"source"`
	SourceRef    string      `json:"sourceRef"`
	Artifacts    interface{} `json:"artifacts"`
	Follow       bool        `json:"follow"`
	CaseTemplate string      `json:"caseTemplate,omitempty"`
	LastSyncDate time.Time   `json:"lastSyncDate,omitempty"`
	Case         string      `json:"case,omitempty"`

	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedBy string    `json:"updatedBy"`
	UpadtedAt time.Time `json:"upadtedAt"`
	User      string    `json:"user,omitempty"`
}

func NewAlert() *Alert {
	var alert *Alert
	alert = new(Alert)
	alert.Severity = Medium
	alert.Date = time.Now()
	alert.TLP = Amber
	alert.Status = New
	alert.Follow = true
	return alert
}

/*
	Alert methods
*/

func (a *Alert) SetType(t string) {
	if t != "" {
		a.Type = t
	}
}

func (a *Alert) SetSource(src string) {
	if src != "" {
		a.Source = src
	}
}

func (a *Alert) SetSourceRef(src string) {
	if src != "" {
		a.SourceRef = src
	}
}

/*
	API Calls
*/

func (api *API) GetAlerts() ([]Alert, error) {
	path := "/api/alert"
	req, err := api.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var alerts []Alert
	var buff []byte
	_, buff, err = api.do(req)

	err = json.Unmarshal(buff, &alerts)
	return alerts, err
}

func (api *API) GetAlert(id string) (Alert, error) {
	if id != "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}

	path := "/api/alert/" + id
	req, err := api.newRequest("GET", path, nil)
	if err != nil {
		return Alert{}, err
	}
	var alert Alert
	var buff []byte
	_, buff, err = api.do(req)

	err = json.Unmarshal(buff, &alert)
	return alert, err
}

func (api *API) SearchAlert(id string) (Alert, error) {
	if id != "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}

	path := "/api/alert/" + id
	req, err := api.newRequest("GET", path, nil)
	if err != nil {
		return Alert{}, err
	}
	var alert Alert
	var buff []byte
	_, buff, err = api.do(req)

	err = json.Unmarshal(buff, &alert)
	return alert, err
}
