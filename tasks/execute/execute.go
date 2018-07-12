// Package execute implements the "execute" task
package execute

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	api "github.com/mesg-foundation/core/api/service"
)

type successOutput struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

type errorOutput struct {
	Message string `json:"message"`
}

type taskInputs struct {
	URL  string `json:"url"`
	Body string `json:"body"`
}

// Task is the "execute" task implementing the task interface
type Task struct {
	HTTPClient http.Client
}

// Process runs one received task and returns key and output for emmision
func (s Task) Process(taskData *api.TaskData) (string, string) {
	var statusCode, body, err = s.doProcess(taskData)
	if err != nil {
		output, _ := json.Marshal(errorOutput{Message: err.Error()})
		return "error", string(output)
	}
	var output, _ = json.Marshal(successOutput{StatusCode: statusCode, Body: string(body)})
	return "success", string(output)
}

func (s Task) doProcess(taskData *api.TaskData) (int, string, error) {
	var ins taskInputs
	if err := json.Unmarshal([]byte(taskData.InputData), &ins); err != nil {
		err = fmt.Errorf("error unmarshalling input of task with ExID [%v]: %v", taskData.ExecutionID, err)
		return 0, "", err
	}

	req, err := http.NewRequest("POST", ins.URL, bytes.NewBuffer([]byte(ins.Body)))
	if err != nil {
		err = fmt.Errorf("error preparing HTTP POST for task with ExID [%v]: %v", taskData.ExecutionID, err)
		return 0, "", err
	}

	postRes, err := s.HTTPClient.Do(req)
	if err != nil {
		err = fmt.Errorf("error running HTTP POST for task with ExID [%v]: %v", taskData.ExecutionID, err)
		return 0, "", err
	}

	body, err := ioutil.ReadAll(postRes.Body)
	if err = postRes.Body.Close(); err != nil {
		err = fmt.Errorf("error closing response payload for task with ExID [%v]: %v", taskData.ExecutionID, err)
		return postRes.StatusCode, string(body), err
	}

	return postRes.StatusCode, string(body), nil
}
