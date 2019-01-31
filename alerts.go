package go4th

import (
	"fmt"

	"gopkg.in/oleiade/reflections.v1"
)

// Alert is the data model for an alert.
type Alert struct {
	ID           string      `json:"id,omitempty"`
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
	Artifacts    []*Artifact `json:"artifacts,omitempty"`
	Follow       bool        `json:"follow,omitempty"`
	CaseTemplate string      `json:"caseTemplate,omitempty"`
	LastSyncDate int64       `json:"lastSyncDate,omitempty"`
	Case         string      `json:"case,omitempty"`
	CreatedBy    string      `json:"createdBy,omitempty"`
	CreatedAt    int64       `json:"createdAt,omitempty"`
	UpdatedBy    string      `json:"updatedBy,omitempty"`
	UpdatedAt    int64       `json:"updatedAt,omitempty"`
	User         string      `json:"user,omitempty"`
}

// NewAlert generates an empty alert with the required fields filled with its defaults
func NewAlert() *Alert {
	var alert *Alert
	alert = new(Alert)
	alert.Severity = Medium
	alert.TLP = Amber
	alert.Status = New
	alert.Follow = true
	return alert
}

var updateKeysAlert = []string{"TLP", "Severity", "Tags",
	"CaseTemplate", "Title", "Description"}

// Fields that are read-only
var keysAlert = []string{"Date", "Status", "Follow", "LastSyncDate",
	"Case", "CreatedBy", "CreatedAt", "UpdatedBy", "User", "ID"}

/*
	Alert methods
*/

// SetTitle sets alert's title. Title couldn't be an empty string, otherwise an error will be returned
func (a *Alert) SetTitle(t string) error {
	if t != "" {
		a.Title = t
		return nil
	}
	return fmt.Errorf("title could not be empty")
}

// SetDescription sets alert's description. Description couldn't be an empty string,
// otherwise an error will be returned
func (a *Alert) SetDescription(d string) error {
	if d != "" {
		a.Description = d
		return nil
	}
	return fmt.Errorf("description could not be empty")
}

// SetType sets alert's type. Type couldn't be an empty string, otherwise an error will be returned
func (a *Alert) SetType(t string) error {
	if t != "" {
		a.Type = t
		return nil
	}
	return fmt.Errorf("type could not be empty")
}

// SetSource sets alert's source. Source couldn't be an empty string, otherwise an error will be returned
func (a *Alert) SetSource(s string) error {
	if s != "" {
		a.Source = s
		return nil
	}
	return fmt.Errorf("source could not be empty")
}

// SetSourceRef sets alert's sourceRef. SourceRef couldn't be an empty string,
// otherwise an error will be returned
func (a *Alert) SetSourceRef(sr string) error {
	if sr != "" {
		a.SourceRef = sr
		return nil
	}
	return fmt.Errorf("sourceRef could not be empty")
}

// AddArtifact adds an artifact to the alert
func (a *Alert) AddArtifact(art *Artifact) {
	a.Artifacts = append(a.Artifacts, art)
}

/*
	API Calls
*/

// GetAlert gets an specific alert. The alert ID must be provided in terms to get the alert.
// If there is an error, an empty Alert will be returned, otherwise the alert is returned with
// nil error.
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

// GetAlerts gets the whole list of alerts. GetAlerts returns a list of Alert or an empty
// list. It can also return an error.
func (api *API) GetAlerts() ([]Alert, error) {
	path := "/api/alert"
	req, err := api.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return api.readResponseAsAlerts(req)
}

// CreateAlert creates an alert. An alert must be provided as parameter it also needs to have
// the required fields filled. Returns the same alert with ID number and same extra information.
// If any error is produced while creating the alert, that error will be returned.
func (api *API) CreateAlert(alert *Alert) (Alert, error) {
	path := "/api/alert"
	req, err := api.newRequest("POST", path, alert)
	if err != nil {
		return Alert{}, err
	}

	return api.readResponseAsAlert(req)
}

// UpdateAlert updates the alert information. The alert ID must me provided as well as
// a map of fields:values that are going to be updated. The fileds couldn't be the ones
// that are readonly and they must be defined in the Alert type. The alert with its fields
// updated is returned, or an empty alert with an error will do it instead
func (api *API) UpdateAlert(id string, values map[string]interface{}) (Alert, error) {
	if id == "" {
		return Alert{}, fmt.Errorf("id must be provided")
	}

	newAlert := NewAlert()
	for field, value := range values {
		if in(field, updateKeysAlert) || !in(field, keysAlert) {
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

// DeleteAlert deletes and alert. The alert ID must be provided. If ID is empty string, an error
// will be returned, otherwise if everything goes well, no error will be returned.
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

// ReadAlert marks an alert as read. The alert ID must be provied and the modified alert is returned.
// If alert ID is empty or there is any other, it is returned.
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

// UnreadAlert marks an alert as unread. The alert ID must be provied and the modified alert is returned.
// If alert ID is empty or there is any other, it is returned.
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

// AlertToCase converts an alert to a case. The alert ID must be provided. If the alert ID is
// empty an error is returned. If everything was ok, the returned alert is the alert converted to case.
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

// FollowAlert switches Follow field to true. The alert ID must be provied otherwise an error
// is returned.
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

// UnfollowAlert switches Follow field to false. The alert ID must be provied otherwise an error
// is returned.
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
