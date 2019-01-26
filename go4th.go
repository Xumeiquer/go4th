package go4th

import (
	"bytes"
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

// NewAPI retuns a new API instance ready to operate
func NewAPI(baseURL, apiKey string) *API {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic("bad base url")
	}
	api := &API{
		baseURL:    u,
		userAgent:  userAgent,
		httpClient: &http.Client{},
		apiKey:     apiKey,
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
		return nil, nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	// err = json.NewDecoder(resp.Body).Decode(v)
	return nil, data, err
}

var userAgent = "go4th/1.0"

type TLP int
type Severity int
type AlertStatus string

const (
	White TLP = 0
	Green TLP = 1
	Amber TLP = 2
	Red   TLP = 3

	Low    Severity = 1
	Medium Severity = 2
	High   Severity = 3

	New      AlertStatus = "New"
	Updated  AlertStatus = "Updated"
	Ignored  AlertStatus = "Ignored"
	Imported AlertStatus = "Imported"
)
