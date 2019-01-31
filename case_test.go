package go4th

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSetTitleCase(t *testing.T) {
	cs := NewCase()

	err := cs.SetTitle("")
	if err == nil {
		t.Errorf("expectig error to be title could not be empty, but no error found")
	}

	err = cs.SetTitle("Test Case")
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}

	if cs.Title != "Test Case" {
		t.Errorf("expecting title to be Test Case, but found %s", cs.Title)
	}
}

func TestSetDescriptionCase(t *testing.T) {
	cs := NewCase()

	err := cs.SetDescription("")
	if err == nil {
		t.Errorf("expectig error to be description could not be empty, but no error found")
	}

	err = cs.SetDescription("Test Case")
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}

	if cs.Description != "Test Case" {
		t.Errorf("expecting description to be Test Case, but found %s", cs.Description)
	}
}

func TestSetSeverityCase(t *testing.T) {
	cs := NewCase()

	err := cs.SetSeverity(Low)
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}

	if cs.Severity != Low {
		t.Errorf("expecting severity to be %d, but found %d", Low, cs.Severity)
	}
}

func TestSetOwnerCase(t *testing.T) {
	cs := NewCase()

	err := cs.SetOwner("")
	if err == nil {
		t.Errorf("expectig error to be owner could not be empty, but no error found")
	}

	err = cs.SetOwner("admin")
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}
	if cs.Owner != "admin" {
		t.Errorf("expectig user to be admin, but found %s", cs.Owner)
	}
}

func TestSetFlagCase(t *testing.T) {
	cs := NewCase()

	err := cs.SetFlag(true)
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}

	err = cs.SetFlag(false)
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}
	if cs.Flag != false {
		t.Errorf("expectig flag to be false, but found %t", cs.Flag)
	}
}

func TestSetTLPCase(t *testing.T) {
	cs := NewCase()

	err := cs.SetTLP(Green)
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}

	err = cs.SetTLP(Red)
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}
	if cs.TLP != Red {
		t.Errorf("expectig TLP to be %d, but found %d", Red, cs.TLP)
	}
}

func TestSetPAPCase(t *testing.T) {
	cs := NewCase()

	err := cs.SetPAP(Green)
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}

	err = cs.SetPAP(Red)
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}
	if cs.PAP != Red {
		t.Errorf("expectig PAP to be %d, but found %d", Red, cs.PAP)
	}
}

func TestSetTagsCase(t *testing.T) {
	cs := NewCase()

	err := cs.SetTags([]string{})
	if err == nil {
		t.Errorf("expectig error to be Tags could not be empty, but no error found")
	}
	tags := []string{"Test", "Case"}
	err = cs.SetTags(tags)
	if err != nil {
		t.Errorf("expectig no error, but found %s", err.Error())
	}

	if len(cs.Tags) != len(tags) {
		t.Errorf("expecting Tags to %d elements, but found %d", len(tags), len(cs.Tags))
	}

	if cs.Tags[0] != "Test" {
		t.Errorf("expecting to be Test, but found %s", cs.Tags[0])
	}
	if cs.Tags[1] != "Case" {
		t.Errorf("expecting to be Test, but found %s", cs.Tags[1])
	}
}

func TestGetCase(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/case/"+globalID {
			t.Errorf("expected path to be /api/case/%s, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		caseCase := NewCase()
		caseCase.ID = globalID

		data, err := json.Marshal(caseCase)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	api := NewAPI(ts.URL, apiKey)
	cas, err := api.GetCase(globalID)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}
	if cas.ID != globalID {
		t.Errorf("expected %s as ID, but found %s", globalID, cas.ID)
	}

	_, err = api.GetCase("")
	if err.Error() != "id must be provided" {
		t.Errorf("expected error to be id must be provided, but found %s", err.Error())
	}
}

func TestGetCases(t *testing.T) {
	caseCase := []Case{Case{}, Case{}}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/case" {
			t.Errorf("expected path to be /api/case, but found %s", r.URL.Path)
		}

		data, err := json.Marshal(caseCase)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	api := NewAPI(ts.URL, apiKey)

	cases, err := api.GetCases()
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if len(cases) != len(caseCase) {
		t.Errorf("expecting %d cases, but found %d", len(caseCase), len(cases))
	}
}

func TestCreateCase(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/case" {
			t.Errorf("expected path to be /api/case, but found %s", r.URL.Path)
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

		var cas *Case
		err := json.NewDecoder(r.Body).Decode(&cas)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}
		cas = NewCase()
		cas.ID = globalID
		cas.CreatedBy = user
		cas.UpdatedBy = user
		cas.CreatedAt = time.Now().Unix()
		cas.UpdatedAt = time.Now().Unix()

		data, err := json.Marshal(cas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	api := NewAPI(ts.URL, apiKey)

	initialCase := newCase()
	cas, err := api.CreateCase(initialCase)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if cas.ID == "" {
		t.Errorf("expected an id, but did not found anything")
	}
	if cas.ID != globalID {
		t.Errorf("expected id to be %s, but found %s", globalID, cas.ID)
	}
	if cas.CreatedBy == "" {
		t.Errorf("expected createdBy, but did not found anything")
	}
	if cas.CreatedBy != user {
		t.Errorf("expected createdBy to be %s, but found %s", user, cas.CreatedBy)
	}
	if cas.UpdatedBy == "" {
		t.Errorf("expected updatedBy, but did not found anything")
	}
	if cas.UpdatedBy != user {
		t.Errorf("expected updatedBy to be %s, but found %s", user, cas.UpdatedBy)
	}
	if cas.CreatedAt == 0 {
		t.Errorf("expected createdAt, but did not found anything")
	}
	if cas.UpdatedAt == 0 {
		t.Errorf("expected updatedAt, but did not found anything")
	}
}

func TestUpdateCase(t *testing.T) {
	ss := httptest.NewServer(newCaseHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			t.Errorf("expected PATCH method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/case/"+globalID {
			t.Errorf("expected path to be /api/case/%s, but found %s", globalID, r.URL.Path)
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

		var cas *Case
		err := json.NewDecoder(r.Body).Decode(&cas)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}
		data, err := json.Marshal(cas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer ts.Close()

	initialCase := newCase()

	api := NewAPI(ss.URL, apiKey)
	cas, _ := api.CreateCase(initialCase)
	api = NewAPI(ts.URL, apiKey)

	up := NewUpdater()
	up.Add("Title", "This has been modified")
	up.Add("Description", "This has been modified")
	newCase, err := api.UpdateCase(cas.ID, up)

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}
	if initialCase.ID != newCase.ID {
		t.Errorf("expected ID to be the same, but it is different")
	}
	if initialCase.Title == newCase.Title {
		t.Errorf("expected different title, but the are the same")
	}
	if initialCase.Description == newCase.Description {
		t.Errorf("expected different title, but the are the same")
	}

	api = NewAPI(ts.URL, apiKey)

	up = NewUpdater()
	up.Add("ID", "This has been modified")
	newCase, err = api.UpdateCase(cas.ID, up)

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}
	if initialCase.ID != newCase.ID {
		t.Errorf("expected ID to be the same, but it is different")
	}

	up = NewUpdater()
	_, err = api.UpdateCase("", up)
	if err == nil {
		t.Errorf("expecting an error, but none found")
	}
	if err.Error() != "id must be provided" {
		t.Errorf("expecting error to be id must be provided, but found %s", err.Error())
	}
}

func TestDeleteCase(t *testing.T) {
	var cas Case
	ss := httptest.NewServer(newCaseHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/case/"+globalID {
			t.Errorf("expected path to be /api/case/%s, but found %s", globalID, r.URL.Path)
		}
		if len(r.Header["Accept"]) >= 1 {
			if r.Header["Accept"][0] != "application/json" {
				t.Errorf("expected Accept to be application/json, but found %s", r.Header["Accept"][0])
			}
		} else {
			t.Errorf("expected at least one Accept header, none was found")
		}

		if cas.ID == "" {
			t.Errorf("extected ID to be %s, but found %s", globalID, cas.ID)
		}
		w.Write([]byte{})
	}))
	defer ts.Close()

	initialCase := newCase()

	api := NewAPI(ss.URL, apiKey)
	cas, _ = api.CreateCase(initialCase)

	api = NewAPI(ts.URL, apiKey)
	err := api.DeleteCase(cas.ID)

	if err != nil {
		t.Errorf("expecting error to be nil, but found %s", err.Error())
	}

	err = api.DeleteCase("")

	if err == nil {
		t.Errorf("expecting an error, but none found")
	}
	if err.Error() != "id must be provided" {
		t.Errorf("expecting error to be id must be provided, but found %s", err.Error())
	}
}

func TestMergeCase(t *testing.T) {
	var casA, casB Case
	portfixCaseA := "asdf"
	portfixCaseB := "ghjk"
	ss := httptest.NewServer(newCaseHandler)
	defer ss.Close()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, but found %s", r.Method)
		}
		if r.URL.Path != "/api/case/"+globalID+portfixCaseA+"/_merge/"+globalID+portfixCaseB {
			t.Errorf("expected path to be /api/case/%s/_merge/%s, but found %s", globalID+portfixCaseA, globalID+portfixCaseB, r.URL.Path)
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
		var data struct {
			CaseID       string `json:"caseId,omitempty"`
			MergedCaseID string `json:"mergedCaseId,omitempty"`
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}
		var res *Case
		res = newCase()
		res.MergeFrom = append(res.MergeFrom, data.CaseID)
		res.MergeFrom = append(res.MergeFrom, data.MergedCaseID)

		resData, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(resData)
	}))
	defer ts.Close()

	caseA := newCase()
	caseB := newCase()

	api := NewAPI(ss.URL, apiKey)
	casA, _ = api.CreateCase(caseA)
	casB, _ = api.CreateCase(caseB)
	// simulate new and different case
	casA.ID += portfixCaseA
	casB.ID += portfixCaseB

	api = NewAPI(ts.URL, apiKey)
	merged, err := api.MergeCase(casA.ID, casB.ID)
	if err != nil {
		t.Errorf("expecting error to be nil, but found %s", err.Error())
	}

	if len(merged.MergeFrom) != 2 {
		t.Errorf("expecting to have two merged case, but found %d", len(merged.MergeFrom))
	}

	if merged.MergeFrom[0] != globalID+portfixCaseA {
		t.Errorf("expecting first case to be %s, but found %s", globalID+portfixCaseA, merged.MergeFrom[0])
	}

	if merged.MergeFrom[1] != (globalID + portfixCaseB) {
		t.Errorf("expecting first case to be %s, but found %s", globalID+portfixCaseB, merged.MergeFrom[0])
	}
}
