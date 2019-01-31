package go4th

import (
	"encoding/json"
	"net/http"
	"time"
)

var globalID = "ff7116c6c1f125a5d50c134139099a618fd54408"
var apiKey = "cd7c97854e86bdb07ebe49a0bd6b4da849a5f8d2"
var user = "admin"
var newAlertHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var alert *Alert
	err := json.NewDecoder(r.Body).Decode(&alert)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	alert = NewAlert()
	alert.ID = globalID
	alert.CreatedBy = user
	alert.UpdatedBy = user
	alert.Date = time.Now().Unix()
	alert.CreatedAt = time.Now().Unix()
	alert.UpdatedAt = time.Now().Unix()
	alert.LastSyncDate = time.Now().Unix()

	data, err := json.Marshal(alert)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
})

var newCaseHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var cas *Case
	err := json.NewDecoder(r.Body).Decode(&cas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cas = NewCase()
	cas.ID = globalID
	cas.Description = "Test Case Description"
	cas.Severity = Low
	cas.StartDate = time.Now().Unix()
	cas.Owner = "admin"
	cas.Flag = false
	cas.TLP = Green
	cas.Tags = []string{}

	data, err := json.Marshal(cas)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
})

func newAlert() *Alert {
	initialAlert := NewAlert()
	initialAlert.SetTitle("Test Alert")
	initialAlert.SetDescription("Desciption for a test alert")
	initialAlert.SetSourceRef("qweqwe")
	initialAlert.SetSource("Golang")
	initialAlert.SetType("alert")
	return initialAlert
}

func newCase() *Case {
	initialCase := NewCase()
	initialCase.SetTitle("Test Case")
	initialCase.SetDescription("Desciption for a test Case")
	initialCase.SetSeverity(Low)
	initialCase.SetOwner("admin")
	initialCase.SetTLP(Green)
	initialCase.SetTags([]string{"one", "two"})
	return initialCase
}
