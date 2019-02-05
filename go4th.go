package go4th

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// API defines the methods to exchenge information between clinet and The Hive
type API struct {
	baseURL    *url.URL
	userAgent  string
	httpClient *http.Client
	apiKey     string
}

// NewAPI returns a new API instance ready to operate with TheHive instance
func NewAPI(baseURL, apiKey string, trustSSL bool) *API {
	var api *API
	u, err := url.Parse(baseURL)
	if err != nil {
		panic("bad base url")
	}
	if apiKey == "" {
		panic("bad apikey")
	}
	if trustSSL {
		api = &API{
			baseURL:    u,
			userAgent:  userAgent,
			httpClient: &http.Client{},
			apiKey:     apiKey,
		}
	} else {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		api = &API{
			baseURL:    u,
			userAgent:  userAgent,
			httpClient: &http.Client{Transport: tr},
			apiKey:     apiKey,
		}
	}
	return api
}

func (api *API) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := api.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", api.userAgent)
	req.Header.Set("Authorization", "Bearer "+api.apiKey)
	return req, nil
}

func (api *API) do(req *http.Request) (*http.Response, []byte, error) {
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return resp, nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	return resp, data, err
}

var userAgent = "go4th/1.0"

// TLP defines the Traffic Light Protocol
type TLP int

// Severity defines the lavels of severity
type Severity int

// Status defines the Status
type Status string

// AlertStatus defines the alert status
type AlertStatus string

// ResolutionStatus defines the case resolution status
type ResolutionStatus string

// ImpactStatus defines the case impact status
type ImpactStatus string

// CaseStatus defines the case status
type CaseStatus string

const (
	// White, Green, Amber, and Red are the accepted TLP values
	White TLP = 0
	Green TLP = 1
	Amber TLP = 2
	Red   TLP = 3

	// Low, Medium, and High are the accepted Severity values
	Low    Severity = 1
	Medium Severity = 2
	High   Severity = 3

	Waiting    Status = "Waiting"
	InProgress Status = "InProgress"
	Completed  Status = "Completed"
	Cancel     Status = "Cancel"

	// New, Updated, Ignored, and Imported are the accepted AlertStatus values
	New      AlertStatus = "New"
	Updated  AlertStatus = "Updated"
	Ignored  AlertStatus = "Ignored"
	Imported AlertStatus = "Imported"

	Indeterminate ResolutionStatus = "Indeterminate"
	FalsePositive ResolutionStatus = "FalsePositive"
	TruePositive  ResolutionStatus = "TruePositive"
	Other         ResolutionStatus = "Other"
	Duplicated    ResolutionStatus = "Duplicated"

	NoImpact      ImpactStatus = "NoImpact"
	WithImpact    ImpactStatus = "WithImpact"
	NotApplicable ImpactStatus = "NotApplicable"

	Open     CaseStatus = "Open"
	Resolved CaseStatus = "Resolved"
	Deleted  CaseStatus = "Deleted"
)

// ApiError represents an error response from The Hive
type ApiError struct {
	TableName string  `json:"tableName,omitempty"`
	Type      string  `json:"type,omitempty"`
	Errors    []Error `json:"errors,omitempty"`
}

// Error is part of the ApiError structure and it conteins a specific error
type Error struct {
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}
