package go4th

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/k0kubun/pp"
)

func TestGetAlert(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		alerts := []Alert{Alert{
			Title:       "Alert test",
			Description: "Alert description",
			Severity:    1,
			Date:        time.Now(),
			Tags:        []string{},
			TLP:         Green,
			Status:      NewAS,
			Type:        "test",
			Source:      "test",
			SourceRef:   "test",
			Artifacts:   nil,
			Follow:      true,
		}}
		data, err := json.Marshal(alerts)

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

	pp.Println(alerts)

}
