package go4th

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSetTitleAlert(t *testing.T) {
	alert := NewAlert()
	title := ""

	err := alert.SetTitle(title)

	if alert.Title != "" {
		t.Errorf("expected title to be empty, but found %s", alert.Title)
	}

	if err.Error() != "title could not be empty" {
		t.Errorf("expected error title could not be empty, but found %s", err.Error())
	}

	title = "test"

	if alert.Title != "" {
		t.Errorf("expected title to be empty, but found %s", alert.Title)
	}

	alert.SetTitle(title)

	if alert.Title != title {
		t.Errorf("expected title to be %s, but found %s", title, alert.Title)
	}
}

func TestSetDescriptionAlert(t *testing.T) {
	alert := NewAlert()
	description := ""

	err := alert.SetDescription(description)

	if alert.Description != "" {
		t.Errorf("expected description to be empty, but found %s", alert.Description)
	}

	if err.Error() != "description could not be empty" {
		t.Errorf("expected error description could not be empty, but found %s", err.Error())
	}

	description = "test"

	if alert.Description != "" {
		t.Errorf("expected description to be empty, but found %s", alert.Description)
	}

	alert.SetDescription(description)

	if alert.Description != description {
		t.Errorf("expected description to be %s, but found %s", description, alert.Description)
	}
}

func TestSetTypeAlert(t *testing.T) {
	alert := NewAlert()
	typ := ""

	err := alert.SetType(typ)

	if alert.Type != "" {
		t.Errorf("expected typ to be empty, but found %s", alert.Type)
	}

	if err.Error() != "type could not be empty" {
		t.Errorf("expected error typ could not be empty, but found %s", err.Error())
	}

	typ = "test"

	if alert.Type != "" {
		t.Errorf("expected typ to be empty, but found %s", alert.Type)
	}

	alert.SetType(typ)

	if alert.Type != typ {
		t.Errorf("expected typ to be %s, but found %s", typ, alert.Type)
	}
}

func TestSetSourceAlert(t *testing.T) {
	alert := NewAlert()
	source := ""

	err := alert.SetSource(source)

	if alert.Source != "" {
		t.Errorf("expected source to be empty, but found %s", alert.Source)
	}

	if err.Error() != "source could not be empty" {
		t.Errorf("expected error source could not be empty, but found %s", err.Error())
	}

	source = "test"

	if alert.Source != "" {
		t.Errorf("expected source to be empty, but found %s", alert.Source)
	}

	alert.SetSource(source)

	if alert.Source != source {
		t.Errorf("expected source to be %s, but found %s", source, alert.Source)
	}
}

func TestSetSourceRefAlert(t *testing.T) {
	alert := NewAlert()
	sourceRef := ""

	err := alert.SetSourceRef(sourceRef)

	if alert.SourceRef != "" {
		t.Errorf("expected sourceRef to be empty, but found %s", alert.SourceRef)
	}

	if err.Error() != "sourceRef could not be empty" {
		t.Errorf("expected error sourceRef could not be empty, but found %s", err.Error())
	}

	sourceRef = "test"

	if alert.SourceRef != "" {
		t.Errorf("expected sourceRef to be empty, but found %s", alert.SourceRef)
	}

	alert.SetSourceRef(sourceRef)

	if alert.SourceRef != sourceRef {
		t.Errorf("expected sourceRef to be %s, but found %s", sourceRef, alert.SourceRef)
	}
}

func TestGetAlert(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert/"+globalID {
			t.Errorf("expected path to be /api/alert/%s, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		alertCase := NewAlert()
		alertCase.ID = globalID

		data, err := json.Marshal(alertCase)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	api := NewAPI(ts.URL, apiKey)
	alert, err := api.GetAlert(globalID)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}
	if alert.ID != globalID {
		t.Errorf("expected %s as ID, but found %s", globalID, alert.ID)
	}

	_, err = api.GetAlert("")
	if err.Error() != "id must be provided" {
		t.Errorf("expected error to be id must be provided, but found %s", err.Error())
	}
}

func TestGetAlerts(t *testing.T) {
	alertsCase := []Alert{Alert{}, Alert{}}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert" {
			t.Errorf("expected path to be /api/alert, but found %s", r.URL.Path)
		}

		data, err := json.Marshal(alertsCase)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	api := NewAPI(ts.URL, apiKey)

	alerts, err := api.GetAlerts()
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if len(alerts) != len(alertsCase) {
		t.Errorf("expecting %d alerts, but found %d", len(alertsCase), len(alerts))
	}
}

func TestCreateAlert(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert" {
			t.Errorf("expected path to be /api/alert, but found %s", r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}
		if len(r.Header["Content-Type"]) >= 1 {
			if r.Header["Content-Type"][0] != "application/json" {
				t.Errorf("expected Content-Type to be application/json, but found %s", r.Header["Content-Type"][0])
			}
		} else {
			t.Errorf("expected at least one Content-Type header, none was found")
		}

		var alert *Alert
		err := json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
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
	}))
	defer ts.Close()

	api := NewAPI(ts.URL, apiKey)

	initialAlert := newAlert()
	alert, err := api.CreateAlert(initialAlert)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if alert.ID == "" {
		t.Errorf("expected an id, but did not found anything")
	}
	if alert.ID != globalID {
		t.Errorf("expected id to be %s, but found %s", globalID, alert.ID)
	}
	if alert.CreatedBy == "" {
		t.Errorf("expected createdBy, but did not found anything")
	}
	if alert.CreatedBy != user {
		t.Errorf("expected createdBy to be %s, but found %s", user, alert.CreatedBy)
	}
	if alert.UpdatedBy == "" {
		t.Errorf("expected updatedBy, but did not found anything")
	}
	if alert.UpdatedBy != user {
		t.Errorf("expected updatedBy to be %s, but found %s", user, alert.UpdatedBy)
	}
	if alert.Date == 0 {
		t.Errorf("expected date, but did not found anything")
	}
	if alert.CreatedAt == 0 {
		t.Errorf("expected createdAt, but did not found anything")
	}
	if alert.UpdatedAt == 0 {
		t.Errorf("expected updatedAt, but did not found anything")
	}
	if alert.LastSyncDate == 0 {
		t.Errorf("expected lastSyncDate, but did not found anything")
	}
}

func TestUpdateAlert(t *testing.T) {
	ss := httptest.NewServer(newAlertHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			t.Errorf("expected PATCH method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert/"+globalID {
			t.Errorf("expected path to be /api/alert/%s, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}
		if len(r.Header["Content-Type"]) >= 1 {
			if r.Header["Content-Type"][0] != "application/json" {
				t.Errorf("expected Content-Type to be application/json, but found %s", r.Header["Content-Type"][0])
			}
		} else {
			t.Errorf("expected at least one Content-Type header, none was found")
		}

		var alert *Alert
		err := json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}
		data, err := json.Marshal(alert)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	initialAlert := newAlert()

	api := NewAPI(ss.URL, apiKey)
	alert, _ := api.CreateAlert(initialAlert)
	api = NewAPI(ts.URL, apiKey)

	up := NewUpdater()
	up.Add("Title", "This has been modified")
	up.Add("Description", "This has been modified")
	newAlert, err := api.UpdateAlert(alert.ID, up)

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}
	if initialAlert.ID != newAlert.ID {
		t.Errorf("expected ID to be the same, but it is different")
	}
	if initialAlert.Title == newAlert.Title {
		t.Errorf("expected different title, but the are the same")
	}
	if initialAlert.Description == newAlert.Description {
		t.Errorf("expected different title, but the are the same")
	}

	api = NewAPI(ts.URL, apiKey)

	up = NewUpdater()
	up.Add("ID", "This has been modified")
	newAlert, err = api.UpdateAlert(alert.ID, up)

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}
	if initialAlert.ID != newAlert.ID {
		t.Errorf("expected ID to be the same, but it is different")
	}

	up = NewUpdater()
	_, err = api.UpdateAlert("", up)
	if err == nil {
		t.Errorf("expecting an error, but none found")
	}
	if err.Error() != "id must be provided" {
		t.Errorf("expecting error to be id must be provided, but found %s", err.Error())
	}
}

func TestDeleteAlert(t *testing.T) {
	var alert Alert
	ss := httptest.NewServer(newAlertHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert/"+globalID {
			t.Errorf("expected path to be /api/alert/%s, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		if alert.ID == "" {
			t.Errorf("extected ID to be %s, but found %s", globalID, alert.ID)
		}
		w.Write([]byte{})
	}))
	defer ts.Close()

	initialAlert := newAlert()

	api := NewAPI(ss.URL, apiKey)
	alert, _ = api.CreateAlert(initialAlert)

	api = NewAPI(ts.URL, apiKey)
	err := api.DeleteAlert(alert.ID)

	if err != nil {
		t.Errorf("expecting error to be nil, but found %s", err.Error())
	}

	err = api.DeleteAlert("")

	if err == nil {
		t.Errorf("expecting an error, but none found")
	}
	if err.Error() != "id must be provided" {
		t.Errorf("expecting error to be id must be provided, but found %s", err.Error())
	}
}

func TestReadAlert(t *testing.T) {
	var alert Alert
	ss := httptest.NewServer(newAlertHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert/"+globalID+"/markAsRead" {
			t.Errorf("expected path to be /api/alert/%s/markAsRead, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		alert.Status = Ignored

		data, err := json.Marshal(alert)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))

	initialAlert := newAlert()

	api := NewAPI(ss.URL, apiKey)
	alert, _ = api.CreateAlert(initialAlert)

	api = NewAPI(ts.URL, apiKey)
	alert, err := api.ReadAlert(alert.ID)
	if err != nil {
		t.Errorf("expecting error to be nil, but found %s", err.Error())
	}
	if initialAlert.Status == alert.Status {
		t.Errorf("expecting to be different from %s, but it is not", alert.Status)
	}
	if alert.Status != Ignored {
		t.Errorf("expecting to be Ignored, but found %s", alert.Status)
	}

	_, err = api.ReadAlert("")

	if err == nil {
		t.Errorf("expecting an error, but none found")
	}
	if err.Error() != "id must be provided" {
		t.Errorf("expecting error to be id must be provided, but found %s", err.Error())
	}
}

func TestUnreadAlert(t *testing.T) {
	var alert Alert
	ss := httptest.NewServer(newAlertHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert/"+globalID+"/markAsUnread" {
			t.Errorf("expected path to be /api/alert/%s/markAsRead, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		alert.Status = New

		data, err := json.Marshal(alert)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	initialAlert := newAlert()

	api := NewAPI(ss.URL, apiKey)
	alert, _ = api.CreateAlert(initialAlert)

	api = NewAPI(ts.URL, apiKey)
	alert, err := api.UnreadAlert(alert.ID)
	if err != nil {
		t.Errorf("expecting error to be nil, but found %s", err.Error())
	}
	if initialAlert.Status != alert.Status {
		t.Errorf("expecting to be different from %s, but it is not", alert.Status)
	}
	if alert.Status != New {
		t.Errorf("expecting to be New, but found %s", alert.Status)
	}

	_, err = api.UnreadAlert("")

	if err == nil {
		t.Errorf("expecting an error, but none found")
	}
	if err.Error() != "id must be provided" {
		t.Errorf("expecting error to be id must be provided, but found %s", err.Error())
	}
}

func TestAlertToCase(t *testing.T) {
	var alert Alert
	ss := httptest.NewServer(newAlertHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert/"+globalID+"/createCase" {
			t.Errorf("expected path to be /api/alert/%s/createCase, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		alert.ID = "13-13-13"

		data, err := json.Marshal(alert)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	initialAlert := newAlert()

	api := NewAPI(ss.URL, apiKey)
	alert, _ = api.CreateAlert(initialAlert)

	api = NewAPI(ts.URL, apiKey)
	newAlert, err := api.AlertToCase(alert.ID)

	if err != nil {
		t.Errorf("unexpecting an error, but found %s", err.Error())
	}
	if newAlert.ID != alert.ID {
		t.Errorf("expecting ID to be different, but it is not")
	}

	_, err = api.AlertToCase("")

	if err == nil {
		t.Errorf("expecting an error, but none found")
	}
	if err.Error() != "id must be provided" {
		t.Errorf("expecting error to be id must be provided, but found %s", err.Error())
	}
}

func TestFollowAlert(t *testing.T) {
	var alert Alert
	ss := httptest.NewServer(newAlertHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert/"+globalID+"/follow" {
			t.Errorf("expected path to be /api/alert/%s/follow, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		alert.Follow = true

		data, err := json.Marshal(alert)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	initialAlert := newAlert()

	api := NewAPI(ss.URL, apiKey)
	alert, _ = api.CreateAlert(initialAlert)

	api = NewAPI(ts.URL, apiKey)
	newAlert, err := api.FollowAlert(alert.ID)
	if err != nil {
		t.Errorf("unexpecting an error, but found %s", err.Error())
	}
	if newAlert.Follow != true {
		t.Errorf("expecting follow to be true, but found %t", newAlert.Follow)
	}

	_, err = api.FollowAlert("")

	if err == nil {
		t.Errorf("expecting an error, but none found")
	}
	if err.Error() != "id must be provided" {
		t.Errorf("expecting error to be id must be provided, but found %s", err.Error())
	}
}

func TestUnfollowAlert(t *testing.T) {
	var alert Alert
	ss := httptest.NewServer(newAlertHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert/"+globalID+"/unfollow" {
			t.Errorf("expected path to be /api/alert/%s/unfollow, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		alert.Follow = false

		data, err := json.Marshal(alert)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	initialAlert := newAlert()

	api := NewAPI(ss.URL, apiKey)
	alert, _ = api.CreateAlert(initialAlert)

	api = NewAPI(ts.URL, apiKey)
	newAlert, err := api.UnfollowAlert(alert.ID)
	if err != nil {
		t.Errorf("unexpecting an error, but found %s", err.Error())
	}
	if newAlert.Follow != false {
		t.Errorf("expecting follow to be false, but found %t", newAlert.Follow)
	}

	_, err = api.UnfollowAlert("")

	if err == nil {
		t.Errorf("expecting an error, but none found")
	}
	if err.Error() != "id must be provided" {
		t.Errorf("expecting error to be id must be provided, but found %s", err.Error())
	}
}
