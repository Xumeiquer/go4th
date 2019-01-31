package go4th

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func in(s string, ls []string) bool {
	for _, b := range ls {
		if b == s {
			return true
		}
	}
	return false
}

// Updater is a map of Alert fields:valeus and it is used for updating alerts
type Updater map[string]interface{}

// NewUpdater retuens a new and empty Updater
func NewUpdater() Updater {
	return Updater{}
}

// Add adds entries into the map
func (u Updater) Add(field string, value interface{}) {
	u[field] = value
}

// Del deletes entries from the map
func (u Updater) Del(field string, value interface{}) {
	delete(u, field)
}

func (api *API) readResponseAsAlert(req *http.Request) (Alert, error) {
	var apiError ApiError
	var alertRes Alert
	var buff []byte
	_, buff, err := api.do(req)

	if err != nil {
		return Alert{}, err
	}

	err = json.Unmarshal(buff, &apiError)
	if err != nil {
		return Alert{}, fmt.Errorf("unable to unmarshal response data as error: %s", err.Error())
	}

	if apiError.TableName == "" && apiError.Type == "" && len(apiError.Errors) == 0 {
		err = json.Unmarshal(buff, &alertRes)
		if err != nil {
			return Alert{}, fmt.Errorf("unable to unmarshal response data: %s", err.Error())
		}
		return alertRes, err
	}

	return Alert{}, getError(apiError)
}

func (api *API) readResponseAsAlerts(req *http.Request) ([]Alert, error) {
	var alertsRes []Alert
	var buff []byte
	_, buff, err := api.do(req)

	if err != nil {
		return []Alert{}, err
	}

	err = json.Unmarshal(buff, &alertsRes)
	if err != nil {
		return []Alert{}, fmt.Errorf("unable to unmarshal response data: %s", err.Error())
	}

	return alertsRes, err
}

func (api *API) readResponseAsCase(req *http.Request) (Case, error) {
	var apiError ApiError
	var caseRes Case
	var buff []byte
	_, buff, err := api.do(req)

	if err != nil {
		return Case{}, err
	}

	err = json.Unmarshal(buff, &apiError)
	if err != nil {
		return Case{}, fmt.Errorf("unable to unmarshal response data as error: %s", err.Error())
	}

	if apiError.TableName == "" && apiError.Type == "" && len(apiError.Errors) == 0 {
		err = json.Unmarshal(buff, &caseRes)
		if err != nil {
			return Case{}, fmt.Errorf("unable to unmarshal response data: %s", err.Error())
		}
		return caseRes, err
	}

	return Case{}, getError(apiError)
}

func (api *API) readResponseAsCases(req *http.Request) ([]Case, error) {
	var casesRes []Case
	var buff []byte
	_, buff, err := api.do(req)

	if err != nil {
		return []Case{}, err
	}

	err = json.Unmarshal(buff, &casesRes)
	if err != nil {
		return []Case{}, fmt.Errorf("unable to unmarshal response data: %s", err.Error())
	}

	return casesRes, err
}

func getError(apiErr ApiError) error {
	if len(apiErr.Errors) != 0 {
		var e string
		e = fmt.Sprintf("%s::", apiErr.Type)
		for _, err := range apiErr.Errors {
			e = fmt.Sprintf("%s%s", e, err.Message)
			break
		}
		return errors.New(e)
	}
	return nil
}
