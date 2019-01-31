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

Go for The Hive is a Golang port of [TheHive4py](https://github.com/TheHive-Project/TheHive4py). This is an API client to comunicate with [TheHive](https://github.com/TheHive-Project/TheHive).

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
)

func main() {
  api := go4th.NewAPI(thehive, apiKey)

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

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|GET         |/api/alert                              |List alerts                           |
|POST        |/api/alert                              |Create an alert                       |
|GET         |/api/alert/:alertId                     |Get an alert                          |
|PATCH       |/api/alert/:alertId                     |Update an alert                       |
|DELETE      |/api/alert/:alertId                     |Delete an alert                       |
|POST        |/api/alert/:alertId/markAsRead          |Mark an alert as read                 |
|POST        |/api/alert/:alertId/markAsUnread        |Mark an alert as unread               |
|POST        |/api/alert/:alertId/createCase          |Create a case from an alert           |
|POST        |/api/alert/:alertId/follow              |Follow an alert                       |
|POST        |/api/alert/:alertId/unfollow            |Unfollow an alert                     |

### Case

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|GET         |/api/case                               |List cases                            |
|POST        |/api/case                               |Create a case                         |
|GET         |/api/case/:caseId                       |Get a case                            |
|PATCH       |/api/case/:caseId                       |Update a case                         |
|DELETE      |/api/case/:caseId                       |Remove a case                         |
|POST        |/api/case/:caseId1/_merge/:caseId2      |Merge two cases                       |

## Missing API calls

### Alert

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|POST        |/api/alert/_search                      |Find alerts                           |
|PATCH       |/api/alert/_bulk                        |Update alerts in bulk                 |
|POST        |/api/alert/_stats                       |Compute stats on alerts               |
|POST        |/api/alert/:alertId/merge/:caseId       |Merge an alert in a case              |

### Case

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|POST        |/api/case/_search                       |Find cases                            |
|PATCH       |/api/case/_bulk                         |Update cases in bulk                  |
|POST        |/api/case/_stats                        |Compute stats on cases                |
|GET         |/api/case/:caseId/links                 |Get list of cases linked to this case |
