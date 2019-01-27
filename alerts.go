package go4th

import (
	"fmt"

	"gopkg.in/oleiade/reflections.v1"
)

type Alert struct {
	ID           string      `json:"id,omitempty"`
	IDdb         string      `json:"_id,omitempty"`
	Typ          string      `json:"_type,omitempty"`
	Patent       string      `json:"_parent,omitempty"`
	Routing      string      `json:"_routing,omitempty"`
	Version      int         `json:"_version,omitempty"`
	Title        string      `json:"title,omitempty"`
	Description  string      `json:"description,omitempty"`
	Severity     Severity    `json:"severity,omitempty"`
	Date         int64       `json:"date,omitempty"`
	Tags         []string    `json:"tags,omitempty"`
	TLP          TLP         `json:"tlp,omitempty"`
	Status       AlertStatus `json:"status,omitempty"`
	Type         string      `json:"type,omitempty"`
	Source       string      `json:"source,omitempty"`
	SourceRef    string      `json:"sourceRef,omitempty"`
	Artifacts    interface{} `json:"artifacts,omitempty"`
	Follow       bool        `json:"follow,omitempty"`
	CaseTemplate string      `json:"caseTemplate,omitempty"`
	LastSyncDate int64       `json:"lastSyncDate,omitempty"`
	Case         string      `json:"case,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	CreatedBy    string      `json:"createdBy,omitempty"`
	CreatedAt    int64       `json:"createdAt,omitempty"`
	UpdatedBy    string      `json:"updatedBy,omitempty"`
	UpdatedAt    int64       `json:"updatedAt,omitempty"`
	User         string      `json:"user,omitempty"`
}

func NewAlert() *Alert {
	var alert *Alert
	alert = new(Alert)
	alert.Severity = Medium
	alert.TLP = Amber
	alert.Status = New
	alert.Follow = true
	return alert
}

var updateKeys = []string{"TLP", "Severity", "Tags", "CaseTemplate", "Title", "Description"}
var keys = []string{"Date", "Status", "Follow", "LastSyncDate", "Case", "CreatedBy", "CreatedAt", "UpdatedBy", "User"}

/*
	Alert methods
*/

func (a *Alert) SetTitle(t string) error {
	if t != "" {
		a.Title = t
		return nil
	}
	return fmt.Errorf("title could not be empty")
}

func (a *Alert) SetDescription(t string) error {
	if t != "" {
		a.Description = t
		return nil
	}
	return fmt.Errorf("description could not be empty")
}

func (a *Alert) SetType(t string) error {
	if t != "" {
		a.Type = t
		return nil
	}
	return fmt.Errorf("type could not be empty")
}

func (a *Alert) SetSource(src string) error {
	if src != "" {
		a.Source = src
		return nil
	}
	return fmt.Errorf("source could not be empty")
}

func (a *Alert) SetSourceRef(src string) error {
	if src != "" {
		a.SourceRef = src
		return nil
	}
	return fmt.Errorf("sourceRef could not be empty")
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

	return api.readResponseAsAlerts(req)
}

func (api *API) GetAlert(id string) (Alert, error) {
	if id == "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}

	path := "/api/alert/" + id
	req, err := api.newRequest("GET", path, nil)
	if err != nil {
		return Alert{}, err
	}

	return api.readResponseAsAlert(req)
}

func (api *API) CreateAlert(alert *Alert) (Alert, error) {
	path := "/api/alert"
	req, err := api.newRequest("POST", path, alert)
	if err != nil {
		return Alert{}, err
	}

	return api.readResponseAsAlert(req)
}

func (api *API) UpdateAlert(id string, values map[string]interface{}) (Alert, error) {
	if id == "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}

	newAlert := NewAlert()
	for field, value := range values {
		if in(field, updateKeys) || !in(field, keys) {
			reflections.SetField(newAlert, field, value)
		}
	}

	path := "/api/alert/" + id
	req, err := api.newRequest("PATCH", path, newAlert)
	if err != nil {
		return Alert{}, err
	}

	return api.readResponseAsAlert(req)
}

func (api *API) DeleteAlert(id string) error {
	if id == "" {
		return fmt.Errorf("id must be provided")
	}
	path := "/api/alert/" + id
	_, err := api.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) ReadAlert(id string) (Alert, error) {
	if id == "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}
	path := "/api/alert/" + id + "/markAsRead"
	req, err := api.newRequest("POST", path, nil)
	if err != nil {
		return Alert{}, err
	}
	return api.readResponseAsAlert(req)
}

func (api *API) UnreadAlert(id string) (Alert, error) {
	if id == "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}
	path := "/api/alert/" + id + "/markAsUnread"
	req, err := api.newRequest("POST", path, nil)
	if err != nil {
		return Alert{}, err
	}
	return api.readResponseAsAlert(req)
}

func (api *API) AlertToCase(id string) (Alert, error) {
	if id == "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}
	path := "/api/alert/" + id + "/createCase"
	req, err := api.newRequest("POST", path, nil)
	if err != nil {
		return Alert{}, err
	}
	return api.readResponseAsAlert(req)
}

func (api *API) FollowAlert(id string) (Alert, error) {
	if id == "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}
	path := "/api/alert/" + id + "/follow"
	req, err := api.newRequest("POST", path, nil)
	if err != nil {
		return Alert{}, err
	}
	return api.readResponseAsAlert(req)
}

func (api *API) UnfollowAlert(id string) (Alert, error) {
	if id == "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}
	path := "/api/alert/" + id + "/unfollow"
	req, err := api.newRequest("POST", path, nil)
	if err != nil {
		return Alert{}, err
	}
	return api.readResponseAsAlert(req)
}
