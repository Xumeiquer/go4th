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

type Updater map[string]interface{}

func NewUpdater() Updater {
	return Updater{}
}
func (u Updater) Add(field string, value interface{}) {
	u[field] = value
}

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
