<p align="center">
  <b>Go for The Hive</b>
</p>
<p align="center">
  <a href="https://travis-ci.com/Xumeiquer/go4th"><img src="https://img.shields.io/travis/com/Xumeiquer/go4th/dev.svg"></a>
  <a href="https://godoc.org/github.com/Xumeiquer/go4th"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"></a>
  <a href="https://goreportcard.com/report/Xumeiquer/go4th"><img src="https://goreportcard.com/badge/github.com/Xumeiquer/go4th"></a>
  <a href="https://opensource.org/licenses/Apache-2.0"><img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"></a>
</p>

---

Go for The Hive is a Golang port of [TheHive4py](https://github.com/TheHive-Project/TheHive4py). This is an API client to communicate with [TheHive](https://github.com/TheHive-Project/TheHive).

## Installation

```
go get github.com/Xumeiquer/go4th
```

## Usage

Go 4 TheHive exposes the whole API through an API object.

```go
package main

import (
  "os"

  "github.com/Xumeiquer/go4th"
)

var (
  thehive = "http://127.0.0.1:9000"
  apiKey  = "apiKey"
  trustSSL = true
)

func main() {
  api := go4th.NewAPI(thehive, apiKey, trustSSL)

  alerts, err := api.GetAlerts()
  if err != nil {
    fmt.Println("error while getting alerts")
    os.Exit(1)
  }

  for _, alert := range alerts {
    fmt.Printf("Got Alert %s with title %s\n", alert.ID, alert.Title)
  }
}
```

## API implementation

### Alert

* [x] List alerts
* [x] Find alerts
* [ ] Update alerts in bulk
* [ ] Compute stats on alerts
* [x] Create an alert
* [x] Get an alert
* [x] Update an alert
* [x] Delete an alert
* [x] Mark an alert as read
* [x] Mark an alert as unread
* [x] Create a case from an alert
* [x] Follow an alert
* [x] Unfollow an alert
* [x] Merge an alert in a case

### Case

* [x] List cases
* [x] Find cases
* [ ] Update cases in bulk
* [ ] Compute stats on cases
* [x] Create a case
* [x] Get a case
* [x] Update a case
* [x] Remove a case
* [ ] Get list of cases linked to this case
* [x] Merge two cases

### Obervable

* [ ] Find observables
* [ ] Compute stats on observables
* [ ] Create an observable
* [ ] Get an observable
* [ ] Remove an observable
* [ ] Update an observable
* [ ] Get list of similar observables
* [ ] Update observables in bulk

### Task

* [ ] Find tasks in a case (deprecated)
* [x] Find tasks
* [ ] Compute stats on tasks
* [x] Get a task
* [x] Update a task
* [x] Create a task

### Log

* [ ] Get logs of the task
* [ ] Find logs in specified task
* [ ] Find logs
* [ ] Create a log
* [ ] Update a log
* [ ] Remove a log
* [ ] Get a log

### User

* [ ] Logout
* [ ] User login
* [ ] Get current user
* [ ] Find user
* [ ] Create a user
* [ ] Get a user
* [ ] Delete a user
* [ ] Update user details
* [ ] Set password
* [ ] Change password
