package go4th

import (
	"fmt"
	"log"
)

// Observable represets an observable
type Observable struct {
	ID         string      `json:"id,omitempty"`
	DataType   string      `json:"dataType,omitempty"`
	CreatedBy  string      `json:"createdBy,omitempty"`
	Sighted    bool        `json:"sighted,omitempty"`
	CreatedAt  int64       `json:"createdAt,omitempty"`
	Tags       []string    `json:"tags,omitempty"`
	Ioc        bool        `json:"ioc,omitempty"`
	Data       string      `json:"data,omitempty"`
	Reports    interface{} `json:"reports,omitempty"`
	TLP        TLP         `json:"tlp,omitempty"`
	Message    string      `json:"message,omitempty"`
	Status     string      `json:"status,omitempty"`
	StartDate  int64       `json:"startDate,omitempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
	Case       *Case       `json:"case,omitempty"`
}

// Attachment represents an attachment for observables
type Attachment struct {
	Name        string   `json:"name,omitempty"`
	Hashes      []string `json:"hashes,omitempty"`
	Size        int      `json:"size,omitempty"`
	ContentType string   `json:"contentType,omitempty"`
	ID          string   `json:"id,omitempty"`
}

// NewCase returns a new case object with default values
func NewObservable() *Observable {
	var o *Observable
	o = new(Observable)
	o.TLP = Amber
	o.Tags = []string{}
	o.Ioc = false
	o.Sighted = false
	return o
}

/*
	Observable methods
*/

// SetDataType sets the observable type
func (o *Observable) SetDataType(dt string) error {
	// TODO: Validate dataType against valid values
	if dt == "" {
		return fmt.Errorf("dataType could not be empty")
	}
	o.DataType = dt
	return nil
}

// SetSighted sets whether the observable is sighted or not
func (o *Observable) SetSighted(sighted bool) error {
	o.Sighted = sighted
	return nil
}

// SetIoc sets whether the observable is a IoC
func (o *Observable) SetIoc(ioc bool) error {
	o.Ioc = ioc
	return nil
}

// SetTags sets a list of tags
func (o *Observable) SetTags(tags []string) error {
	if len(tags) == 0 {
		return fmt.Errorf("tags could not be empty")
	}
	o.Tags = tags
	return nil
}

// AddTag adds one tag to the tag list
func (o *Observable) AddTag(tag string) error {
	if tag == "" {
		return fmt.Errorf("tag could not be empty")
	}
	for _, t := range o.Tags {
		if t == tag {
			return fmt.Errorf("tag already in tags")
		}
	}
	o.Tags = append(o.Tags, tag)
	return nil
}

// RemoveTag remove a tag from the tag list
func (o *Observable) RemoveTag(tag string) error {
	if tag == "" {
		return fmt.Errorf("tag could not be empty")
	}
	tagFound := true
	for idx, t := range o.Tags {
		if t == tag {
			if idx < len(o.Tags) {
				o.Tags = append(o.Tags[:idx], o.Tags[idx+1:]...)
			} else {
				o.Tags = append(o.Tags[:idx])
			}
		} else {
			tagFound = false
		}
	}
	if !tagFound {
		return fmt.Errorf("tag not found")
	}
	return nil
}

// SetData sets the observable data
// If you need to provide multiple observable you have to instanciate several Observable objects
// one per Observable.
func (o *Observable) SetData(data string) error {
	if data == "" {
		return fmt.Errorf("data could not be empty")
	}
	o.Data = data
	return nil
}

// SetTLP sets the observable TLP
func (o *Observable) SetTLP(tlp TLP) error {
	o.TLP = tlp
	return nil
}

// SetMessage sets the observable description message
func (o *Observable) SetMessage(message string) error {
	if message == "" {
		return fmt.Errorf("message could not be empty")
	}
	o.Message = message
	return nil
}

/*
	API Calls
*/

// GetObservables get a whole list of observables
func (api *API) GetObservables() ([]Observable, error) {
	path := "/api/case/artifact/_search"
	req, err := api.newRequest("POST", path, nil)
	if err != nil {
		return []Observable{}, err
	}

	return api.readResponseAsObservables(req)
}

// GetObservable get a observable based on its ID
func (api *API) GetObservable(id string) (*Observable, error) {
	if id == "" {
		return &Observable{}, fmt.Errorf("id could not be empty")
	}
	path := "/api/case/artifact/"
	req, err := api.newRequest("GET", path+id, nil)
	if err != nil {
		return &Observable{}, err
	}

	return api.readResponseAsObservable(req)
}

// GetObservableStats gets an observable stats
func (api *API) GetObservableStats(stats *Stats) (ObservableStats, error) {
	path := "/api/case/artifact/_stats"
	req, err := api.newRequest("POST", path, stats)
	if err != nil {
		return ObservableStats{}, err
	}

	return api.readResponseAsObservablesStats(req)
}

// CreateObservable create an observable associated to a case
func (api *API) CreateObservable(caseId string, observable *Observable) (*Observable, error) {
	if caseId == "" {
		return &Observable{}, fmt.Errorf("caseId could not be empty")
	}
	path := "/api/case/%s/artifact"
	req, err := api.newRequest("POST", fmt.Sprintf(path, caseId), observable)
	if err != nil {
		return &Observable{}, err
	}

	return api.readResponseAsObservable(req)
}

// DeleteObservable deletes an observable based on its ID
func (api *API) DeleteObservable(id string) error {
	if id == "" {
		return fmt.Errorf("id could not be empty")
	}
	path := "/api/case/artifact/%s"
	req, err := api.newRequest("DELETE", fmt.Sprintf(path, id), nil)
	if err != nil {
		return err
	}
	_, _, err = api.do(req)
	if err != nil {
		return err
	}
	return nil
}

// UpdateObservable [NotImplemented]
func (api *API) UpdateObservable(observable *Observable) (*Observable, error) {
	log.Println("Not Implemented")
	return nil, nil
}

// GetSimilarObservable gets a list of similar observables based on an observable ID
func (api *API) GetSimilarObservable(id string) ([]Observable, error) {
	if id == "" {
		return []Observable{}, fmt.Errorf("id could not be empty")
	}
	path := "/api/case/artifact/%s/similar"
	req, err := api.newRequest("GET", fmt.Sprintf(path, id), nil)
	if err != nil {
		return []Observable{}, err
	}

	return api.readResponseAsObservables(req)
}
