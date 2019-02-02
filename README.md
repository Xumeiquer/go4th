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
|POST        |/api/alert/_search                      |Find alerts                           |
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
|POST        |/api/case/_search                       |Find cases                            |
|GET         |/api/case/:caseId                       |Get a case                            |
|PATCH       |/api/case/:caseId                       |Update a case                         |
|DELETE      |/api/case/:caseId                       |Remove a case                         |
|POST        |/api/case/:caseId1/_merge/:caseId2      |Merge two cases                       |

### Task

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|GET         |/api/case/task/:taskId                  |Get a task                            |
|POST        |/api/case/:caseId/task                  |Create a task                         |
|PATCH       |/api/case/task/:taskId                  |Update a task                         |
|POST        |/api/case/task/_search                  |Find tasks                            |

## Missing API calls

### Alert

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|PATCH       |/api/alert/_bulk                        |Update alerts in bulk                 |
|POST        |/api/alert/_stats                       |Compute stats on alerts               |
|POST        |/api/alert/:alertId/merge/:caseId       |Merge an alert in a case              |

### Case

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|PATCH       |/api/case/_bulk                         |Update cases in bulk                  |
|POST        |/api/case/_stats                        |Compute stats on cases                |
|GET         |/api/case/:caseId/links                 |Get list of cases linked to this case |

### Task

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|POST        |/api/case/:caseId/task/_search          |Find tasks in a case (deprecated)     |
|POST        |/api/case/task/_stats                   |Compute stats on tasks                |

### Artifact

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|POST        |/api/case/artifact/_search              |Find observables                      |
|POST        |/api/case/artifact/_stats               |Compute stats on observables          |
|POST        |/api/case/:caseId/artifact              |Create an observable                  |
|GET         |/api/case/artifact/:artifactId          |Get an observable                     |
|DELETE      |/api/case/artifact/:artifactId          |Remove an observable                  |
|PATCH       |/api/case/artifact/:artifactId          |Update an observable                  |
|GET         |/api/case/artifact/:artifactId/similar  |Get list of similar observables       |
|PATCH       |/api/case/artifact/_bulk                |Update observables in bulk            |

### Log

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|GET         |/api/case/task/:taskId/log              |Get logs of the task                  |
|POST        |/api/case/task/:taskId/log/_search      |Find logs in specified task           |
|POST        |/api/case/task/log/_search              |Find logs                             |
|POST        |/api/case/task/:taskId/log              |Create a log                          |
|PATCH       |/api/case/task/log/:logId               |Update a log                          |
|DELETE      |/api/case/task/log/:logId               |Remove a log                          |
|GET         |/api/case/task/log/:logId               |Get a log                             |

### User

|HTTP Method |URI                                     |Action                                |
|------------|----------------------------------------|--------------------------------------|
|GET         |/api/logout                             |Logout                                |
|POST        |/api/login                              |User login                            |
|GET         |/api/user/current                       |Get current user                      |
|POST        |/api/user/_search                       |Find user                             |
|POST        |/api/user                               |Create a user                         |
|GET         |/api/user/:userId                       |Get a user                            |
|DELETE      |/api/user/:userId                       |Delete a user                         |
|PATCH       |/api/user/:userId                       |Update user details                   |
|POST        |/api/user/:userId/password/set          |Set password                          |
|POST        |/api/user/:userId/password/change       |Change password                       |