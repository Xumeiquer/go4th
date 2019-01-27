package go4th

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSetTitle(t *testing.T) {
	alert := NewAlert()
	title := "test"

	if alert.Title != "" {
		t.Errorf("expected title to be empty, but found %s", alert.Title)
	}

	alert.SetTitle(title)

	if alert.Title != title {
		t.Errorf("expected title to be %s, but found %s", title, alert.Title)
	}
}

func TestSetTitleEmpty(t *testing.T) {
	alert := NewAlert()
	title := ""

	err := alert.SetTitle(title)

	if alert.Title != "" {
		t.Errorf("expected title to be empty, but found %s", alert.Title)
	}

	if err.Error() != "title could not be empty" {
		t.Errorf("expected error title could not be empty, but found %s", err.Error())
	}
}

func TestSetDescription(t *testing.T) {
	alert := NewAlert()
	description := "test"

	if alert.Description != "" {
		t.Errorf("expected description to be empty, but found %s", alert.Description)
	}

	alert.SetDescription(description)

	if alert.Description != description {
		t.Errorf("expected description to be %s, but found %s", description, alert.Description)
	}
}

func TestSetDescriptionEmpty(t *testing.T) {
	alert := NewAlert()
	description := ""

	err := alert.SetDescription(description)

	if alert.Description != "" {
		t.Errorf("expected description to be empty, but found %s", alert.Description)
	}

	if err.Error() != "description could not be empty" {
		t.Errorf("expected error description could not be empty, but found %s", err.Error())
	}
}

func TestSetType(t *testing.T) {
	alert := NewAlert()
	typ := "test"

	if alert.Type != "" {
		t.Errorf("expected typ to be empty, but found %s", alert.Type)
	}

	alert.SetType(typ)

	if alert.Type != typ {
		t.Errorf("expected typ to be %s, but found %s", typ, alert.Type)
	}
}

func TestSetTypeEmpty(t *testing.T) {
	alert := NewAlert()
	typ := ""

	err := alert.SetType(typ)

	if alert.Type != "" {
		t.Errorf("expected typ to be empty, but found %s", alert.Type)
	}

	if err.Error() != "type could not be empty" {
		t.Errorf("expected error typ could not be empty, but found %s", err.Error())
	}
}

func TestSetSource(t *testing.T) {
	alert := NewAlert()
	source := "test"

	if alert.Source != "" {
		t.Errorf("expected source to be empty, but found %s", alert.Source)
	}

	alert.SetSource(source)

	if alert.Source != source {
		t.Errorf("expected source to be %s, but found %s", source, alert.Source)
	}
}

func TestSetSourceEmpty(t *testing.T) {
	alert := NewAlert()
	source := ""

	err := alert.SetSource(source)

	if alert.Source != "" {
		t.Errorf("expected source to be empty, but found %s", alert.Source)
	}

	if err.Error() != "source could not be empty" {
		t.Errorf("expected error source could not be empty, but found %s", err.Error())
	}
}

func TestSetSourceRef(t *testing.T) {
	alert := NewAlert()
	sourceRef := "test"

	if alert.SourceRef != "" {
		t.Errorf("expected sourceRef to be empty, but found %s", alert.SourceRef)
	}

	alert.SetSourceRef(sourceRef)

	if alert.SourceRef != sourceRef {
		t.Errorf("expected sourceRef to be %s, but found %s", sourceRef, alert.SourceRef)
	}
}

func TestSetSourceRefEmpty(t *testing.T) {
	alert := NewAlert()
	sourceRef := ""

	err := alert.SetSourceRef(sourceRef)

	if alert.SourceRef != "" {
		t.Errorf("expected sourceRef to be empty, but found %s", alert.SourceRef)
	}

	if err.Error() != "sourceRef could not be empty" {
		t.Errorf("expected error sourceRef could not be empty, but found %s", err.Error())
	}
}

func TestGetAlerts(t *testing.T) {
	alertsCase := []Alert{Alert{
		Title:       "Alert test",
		Description: "Alert description",
		Severity:    1,
		Date:        time.Now().Unix(),
		Tags:        []string{},
		TLP:         Green,
		Status:      New,
		Type:        "test",
		Source:      "test",
		SourceRef:   "test",
		Artifacts:   nil,
		Follow:      true,
	}}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET method, but found %s", r.Method)
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
		data, err := json.Marshal(alertsCase)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	api := NewAPI(ts.URL, "123-123-123")

	alerts, err := api.GetAlerts()
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if len(alerts) != 1 {
		t.Errorf("expecting %d alerts, but found %d", len(alertsCase), len(alerts))
	}
}

func TestGetAlert(t *testing.T) {
	ID := "123-123-123-123"
	alertCase := Alert{
		ID:          ID,
		Title:       "Alert test",
		Description: "Alert description",
		Severity:    1,
		Date:        time.Now().Unix(),
		Tags:        []string{},
		TLP:         Green,
		Status:      New,
		Type:        "test",
		Source:      "test",
		SourceRef:   "test",
		Artifacts:   nil,
		Follow:      true,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/alert/123-123-123-123" {
			t.Errorf("expected path to be /api/alert/123-123-123-123, but found %s", r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		data, err := json.Marshal(alertCase)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	api := NewAPI(ts.URL, "123-123-123")
	alert, err := api.GetAlert(ID)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if alert.ID != ID {
		t.Errorf("expected %s as ID, but found %s", ID, alert.ID)
	}
}

func TestGetAlertEmpyID(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	api := NewAPI(ts.URL, "123-123-123")
	_, err := api.GetAlert("")
	if err.Error() != "id must be provided" {
		t.Errorf("expected error to be id must be provided, but found %s", err.Error())
	}
}

func TestCreateAlert(t *testing.T) {
	alertCase := NewAlert()
	alertCase.SetTitle("Test Alert")
	alertCase.SetDescription("Desciption for a test alert")
	alertCase.SetSourceRef("qweqwe")
	alertCase.SetSource("Golang")
	alertCase.SetType("alert")

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
		var alert Alert
		err := json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}

		if alert.Title != "Test Alert" {
			t.Errorf("expecting alert title to be Test Alert, but found %s", alert.Title)
		}
		if alert.Description != "Desciption for a test alert" {
			t.Errorf("expecting alert description to be Desciption for a test alert, but found %s", alert.Description)
		}
		if alert.SourceRef != "qweqwe" {
			t.Errorf("expecting alert sourceRef to be qweqwe, but found %s", alert.SourceRef)
		}
		if alert.Source != "Golang" {
			t.Errorf("expecting alert source to be Golang, but found %s", alert.Source)
		}
		if alert.Type != "alert" {
			t.Errorf("expecting alert type to be alert, but found %s", alert.Type)
		}

		alert.ID = "123-123-123"
		alert.CreatedBy = "admin"
		alert.UpdatedBy = "admin"
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

	api := NewAPI(ts.URL, "123-123-123")
	alert, err := api.CreateAlert(alertCase)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if alert.ID == "" {
		t.Errorf("expected an id, but did not found anything")
	}
	if alert.CreatedBy == "" {
		t.Errorf("expected createdBy, but did not found anything")
	}
	if alert.UpdatedBy == "" {
		t.Errorf("expected updatedBy, but did not found anything")
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
