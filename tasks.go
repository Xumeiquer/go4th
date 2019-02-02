package go4th

import (
	"fmt"

	"gopkg.in/oleiade/reflections.v1"
)

// Task represents a Task
type Task struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      Status `json:"status,omitempty"`
	Flag        bool   `json:"flag,omitempty"`
	Group       string `json:"group,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Order       int    `json:"order,omitempty"`
	StartDate   int64  `json:"startDate,omitempty"`
	EndDate     int64  `json:"endDate,omitempty"`
	CreatedBy   string `json:"createdBy,omitempty"`
	CreatedAt   int64  `json:"createdAt,omitempty"`
	UpdatedBy   string `json:"updatedBy,omitempty"`
	UpdatedAt   int64  `json:"updatedAt,omitempty"`
	User        string `json:"user,omitempty"`
}

// NewTask will return a new task with default values defined
func NewTask() *Task {
	var t *Task
	t = new(Task)
	t.Status = Waiting
	t.Flag = false
	return t
}

var updateKeysTask = []string{"Title", "Description", "Status", "Order",
	"User", "Owner", "Flag", "EndDate"}

/*
	Case methods
*/

// SetTitle sets task's title
func (t *Task) SetTitle(title string) error {
	if title == "" {
		return fmt.Errorf("title could not be empty")
	}
	t.Title = title
	return nil
}

// SetStatus sets task's status
func (t *Task) SetStatus(s Status) error {
	t.Status = s
	return nil
}

// SetFlag sets task's flag
func (t *Task) SetFlag(f bool) error {
	t.Flag = f
	return nil
}

// SetOwner sets task's owner
func (t *Task) SetOwner(owner string) error {
	if owner == "" {
		return fmt.Errorf("owner could not be empty")
	}
	t.Owner = owner
	return nil
}

// SetDescription sets task's description
func (t *Task) SetDescription(description string) error {
	if description == "" {
		return fmt.Errorf("description could not be empty")
	}
	t.Description = description
	return nil
}

// SetGroup sets task's group
func (t *Task) SetGroup(group string) error {
	if group == "" {
		return fmt.Errorf("group could not be empty")
	}
	t.Group = group
	return nil
}

/*
	API Calls
*/

// GetTask gets task based on its ID
func (api *API) GetTask(id string) (Task, error) {
	if id == "" {
		return Task{}, fmt.Errorf("id must be provided")
	}

	path := "/api/case/task/" + id
	req, err := api.newRequest("GET", path, nil)
	if err != nil {
		return Task{}, err
	}

	return api.readResponseAsTask(req)
}

// CreateTask created a task associated to an case ID
func (api *API) CreateTask(caseID string, task *Task) (Task, error) {
	if caseID == "" {
		return Task{}, fmt.Errorf("caseId could not be empty")
	}

	path := "/api/case/" + caseID + "/task"
	req, err := api.newRequest("POST", path, task)
	if err != nil {
		return Task{}, err
	}

	return api.readResponseAsTask(req)
}

// UpdateTask updates a task based on its ID
func (api *API) UpdateTask(id string, values map[string]interface{}) (Task, error) {
	if id == "" {
		return Task{}, fmt.Errorf("id must be provided")
	}

	newTask := NewTask()
	for field, value := range values {
		if in(field, updateKeysTask) {
			reflections.SetField(newTask, field, value)
		}
	}

	path := "/api/case/task/" + id
	req, err := api.newRequest("PATCH", path, newTask)
	if err != nil {
		return Task{}, err
	}

	return api.readResponseAsTask(req)
}

// SearchTask searches tasks based on the query
func (api *API) SearchTask(query *Query) ([]Task, error) {
	path := "/api/case/task/_search"
	req, err := api.newRequest("POST", path, query)
	if err != nil {
		return []Task{}, err
	}
	return api.readResponseAsTasks(req)
}
