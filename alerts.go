package go4th

import (
	"encoding/json"
	"fmt"
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
	var alerts []Alert
	var buff []byte
	_, buff, err = api.do(req)

	err = json.Unmarshal(buff, &alerts)
	return alerts, err
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
	var alert Alert
	var buff []byte
	_, buff, err = api.do(req)

	err = json.Unmarshal(buff, &alert)
	return alert, err
}

func (api *API) CreateAlert(alert *Alert) (Alert, error) {
	path := "/api/alert"
	req, err := api.newRequest("POST", path, alert)
	if err != nil {
		return Alert{}, err
	}

	var alertRes Alert
	var buff []byte
	_, buff, err = api.do(req)

	err = json.Unmarshal(buff, &alertRes)
	if err != nil {
		var apiError ApiError
		err = json.Unmarshal(buff, &apiError)
		if err != nil {
			return Alert{}, err
		}
		return Alert{}, fmt.Errorf("%s::%s", apiError.Type, apiError.Message)
	}
	return alertRes, err
}

func (api *API) UpdateAlert(id string, alert *Alert, fields []string) (Alert, error) {
	if id == "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}

	newAlert := NewAlert()
	for _, field := range fields {
		if in(field, updateKeys) {
			switch field {
			case "TLP":
				if alert.TLP != 0 {
					newAlert.TLP = alert.TLP
				}
			case "Severity":
				if alert.Severity != 0 {
					newAlert.Severity = alert.Severity
				}
			case "Tags":
				if len(alert.Tags) != 0 {
					newAlert.Tags = alert.Tags
				}
			case "CaseTemplate":
				if alert.CaseTemplate != "" {
					newAlert.CaseTemplate = alert.CaseTemplate
				}
			case "Title":
				if alert.Title != "" {
					newAlert.Title = alert.Title
				}
			case "Description":
				if alert.Description != "" {
					newAlert.Description = alert.Description
				}
			}
		}
	}

	for _, field := range fields {
		if in(field, keys) {
			switch field {
			case "Date":
				if alert.Date != 0 {
					newAlert.Date = alert.Date
				}
			case "Status":
				if alert.Status != "" {
					newAlert.Status = alert.Status
				}
			case "Follow":
				if alert.Follow != newAlert.Follow {
					newAlert.Follow = alert.Follow
				}
			case "LastSyncDate":
				if alert.LastSyncDate != 0 {
					newAlert.LastSyncDate = alert.LastSyncDate
				}
			case "Case":
				if alert.Case != "" {
					newAlert.Case = alert.Case
				}
			case "CreatedBy":
				if alert.CreatedBy != "" {
					newAlert.CreatedBy = alert.CreatedBy
				}
			case "CreatedAt":
				if alert.CreatedAt != 0 {
					newAlert.CreatedAt = alert.CreatedAt
				}
			case "UpdatedBy":
				if alert.UpdatedBy != "" {
					newAlert.UpdatedBy = alert.UpdatedBy
				}
			case "User":
				if alert.User != "" {
					newAlert.User = alert.User
				}
			}
		}
	}

	path := "/api/alert/" + id
	req, err := api.newRequest("PATCH", path, newAlert)
	if err != nil {
		return Alert{}, err
	}

	var alertRes Alert
	var buff []byte
	_, buff, err = api.do(req)

	err = json.Unmarshal(buff, &alertRes)
	if err != nil {
		var apiError ApiError
		err = json.Unmarshal(buff, &apiError)
		if err != nil {
			return Alert{}, err
		}
		return Alert{}, fmt.Errorf("%s::%s", apiError.Type, apiError.Message)
	}
	return alertRes, err
}
