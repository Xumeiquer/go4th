package go4th

import (
	"encoding/json"
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
	var alertRes Alert
	var buff []byte
	_, buff, err := api.do(req)

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

func (api *API) readResponseAsAlerts(req *http.Request) ([]Alert, error) {
	var alerts []Alert
	var buff []byte
	_, buff, err := api.do(req)

	err = json.Unmarshal(buff, &alerts)
	return alerts, err
}
