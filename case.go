package go4th

import (
	"fmt"
	"time"

	"gopkg.in/oleiade/reflections.v1"
)

// Case is the data model for an Cases.
type Case struct {
	ID               string           `json:"id,omitempty"`
	Title            string           `json:"title,omitempty"`
	Description      string           `json:"description,omitempty"`
	Severity         Severity         `json:"severity,omitempty"`
	StartDate        int64            `json:"startDate,omitempty"`
	Owner            string           `json:"owner,omitempty"`
	Flag             bool             `json:"flag,omitempty"`
	TLP              TLP              `json:"tlp,omitempty"`
	PAP              TLP              `json:"pap,omitempty"`
	Tags             []string         `json:"tags,omitempty"`
	ResolutionStatus ResolutionStatus `json:"resolutionStatus,omitempty"`
	ImpactStatus     ImpactStatus     `json:"impactStatus,omitempty"`
	Summary          string           `json:"summary,omitempty"`
	EndDate          int64            `json:"endDate,omitempty"`
	Metrics          interface{}      `json:"metrics,omitempty"`
	Status           CaseStatus       `json:"status,omitempty"`
	CaseID           int              `json:"caseID,omitempty"`
	MergeInto        string           `json:"mergeInto,omitempty"`
	MergeFrom        []string         `json:"mergeFrom,omitempty"`
	CustomField      CustomField      `json:"customFields,omitempty"`
	CreatedBy        string           `json:"createdBy,omitempty"`
	CreatedAt        int64            `json:"createdAt,omitempty"`
	UpdatedBy        string           `json:"updatedBy,omitempty"`
	UpdatedAt        int64            `json:"updatedAt,omitempty"`
	User             string           `json:"user,omitempty"`
}

// NewCase returns a new case object with default values
func NewCase() *Case {
	var c *Case
	c = new(Case)
	c.Severity = Medium
	c.StartDate = time.Now().Unix()
	c.Flag = false
	c.TLP = Amber
	c.Tags = []string{}
	c.Status = Open
	return c
}

var updateKeysCase = []string{"Title", "Description", "Severity", "StartDate", "Owner", "Flag", "TLP",
	"Tags", "Status", "ResolutionStatus", "ImpactStatus", "Summary", "EndDate", "Metrics", "CustomFields"}

// Fields that are read-only
var keysCase = []string{"Date", "Status", "Follow", "LastSyncDate",
	"Case", "CreatedBy", "CreatedAt", "UpdatedBy", "User", "ID"}

/*
	Case methods
*/

// SetTitle sets Case's title
func (c *Case) SetTitle(title string) error {
	if title == "" {
		return fmt.Errorf("title could not be empty")
	}
	c.Title = title
	return nil
}

// SetDescription sets Case's description
func (c *Case) SetDescription(description string) error {
	if description == "" {
		return fmt.Errorf("Description could not be empty")
	}
	c.Description = description
	return nil
}

func (c *Case) SetStatus(s CaseStatus) error {
	c.Status = s
	return nil
}

// SetSeverity sets Case's severity
func (c *Case) SetSeverity(severity Severity) error {
	c.Severity = severity
	return nil
}

// SetOwner sets Case's owner
func (c *Case) SetOwner(owner string) error {
	if owner == "" {
		return fmt.Errorf("Owner could not be empty")
	}
	c.Owner = owner
	return nil
}

// SetFlag sets Case's Flag
func (c *Case) SetFlag(flag bool) error {
	c.Flag = flag
	return nil
}

// SetTLP sets Case's TLP
func (c *Case) SetTLP(tlp TLP) error {
	c.TLP = tlp
	return nil
}

// SetPAP sets Case's PAP
func (c *Case) SetPAP(pap TLP) error {
	c.PAP = pap
	return nil
}

// SetTags sets Case's Tags
func (c *Case) SetTags(tags []string) error {
	if len(tags) == 0 {
		return fmt.Errorf("Tags could not be empty")
	}
	c.Tags = tags
	return nil
}

/*
	API Calls
*/

// GetCase gets an specific case. The case ID must be provided in terms to get the case.
// If there is an error, an empty case will be returned, otherwise the case is returned with
// nil error.
func (api *API) GetCase(id string) (Case, error) {
	if id == "" {
		return Case{}, fmt.Errorf("id must be provided")
	}

	path := "/api/case/" + id
	req, err := api.newRequest("GET", path, nil)
	if err != nil {
		return Case{}, err
	}

	return api.readResponseAsCase(req)
}

// GetCases gets the whole list of cases. GetCases returns a list of Alert or an empty
// list. It can also return an error.
func (api *API) GetCases() ([]Case, error) {
	path := "/api/case"
	req, err := api.newRequest("GET", path, nil)
	if err != nil {
		return []Case{}, err
	}

	return api.readResponseAsCases(req)
}

// CreateCase creates an case. An case must be provided as parameter it also needs to have
// the required fields filled. Returns the same case with ID number and same extra information.
// If any error is produced while creating the case, that error will be returned.
func (api *API) CreateCase(cas *Case) (Case, error) {
	path := "/api/case"
	req, err := api.newRequest("POST", path, cas)
	if err != nil {
		return Case{}, err
	}

	return api.readResponseAsCase(req)
}

// UpdateCase updates the case information. The case ID must me provided as well as
// a map of fields:values that are going to be updated. The fileds couldn't be the ones
// that are readonly and they must be defined in the Case type. The case with its fields
// updated is returned, or an empty case with an error will do it instead
func (api *API) UpdateCase(id string, values map[string]interface{}) (Case, error) {
	if id == "" {
		return Case{}, fmt.Errorf("id must be provided")
	}

	newCase := NewCase()
	for field, value := range values {
		if in(field, updateKeysCase) || !in(field, keysCase) {
			reflections.SetField(newCase, field, value)
		}
	}

	path := "/api/case/" + id
	req, err := api.newRequest("PATCH", path, newCase)
	if err != nil {
		return Case{}, err
	}

	return api.readResponseAsCase(req)
}

// DeleteCase deletes and case. The case ID must be provided. If ID is empty string, an error
// will be returned, otherwise if everything goes well, no error will be returned.
func (api *API) DeleteCase(id string) error {
	if id == "" {
		return fmt.Errorf("id must be provided")
	}
	path := "/api/case/" + id
	_, err := api.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

// MergeCase merges one case into another. Both ID must be provided, otherwise an error will be
// returned. If everything goes well, MegeCase will return a merged case.
func (api *API) MergeCase(id, mergeID string) (Case, error) {
	if id == "" {
		return Case{}, fmt.Errorf("id must be provided")
	}
	if mergeID == "" {
		return Case{}, fmt.Errorf("mergeCaseId must be provided")
	}

	data := struct {
		CaseID       string `json:"caseId,omitempty"`
		MergedCaseID string `json:"mergedCaseId,omitempty"`
	}{
		CaseID:       id,
		MergedCaseID: mergeID,
	}
	path := "/api/case/" + id + "/_merge/" + mergeID
	req, err := api.newRequest("POST", path, data)
	if err != nil {
		return Case{}, err
	}

	return api.readResponseAsCase(req)
}

// SearchCase searches cases based on the query
func (api *API) SearchCase(query *Query) ([]Case, error) {

	path := "/api/case/_search"
	req, err := api.newRequest("POST", path, query)
	if err != nil {
		return []Case{}, err
	}
	return api.readResponseAsCases(req)
}
