// Package execute_many implements the "executeMany" task
package execute_many

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	api "github.com/mesg-foundation/core/api/service"
)

type successOutput struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

type errorOutput struct {
	Message string `json:"message"`
}

type responseOutput struct {
	Success successOutput `json:"success"`
	Error   errorOutput   `json:"error"`
}

type requestsInput struct {
	URL  string `json:"url"`
	Body string `json:"body"`
}

type taskInputs struct {
	Requests []requestsInput `json:"requests"`
	Async    bool            `json:"async"`
}

// Task is the "executeMany" task implementing the task interface
type Task struct {
	HTTPClient http.Client
}

// Process runs one received task and returns key and output for emision
func (s Task) Process(taskData *api.TaskData) (string, string) {
	var ins taskInputs
	if err := json.Unmarshal([]byte(taskData.InputData), &ins); err != nil {
		err = fmt.Errorf("error unmarshalling input of task with ExID [%v]: %v", taskData.ExecutionID, err)
		bytOutput, _ := json.Marshal(errorOutput{err.Error()})
		return "error", string(bytOutput)
	}

	var (
		responses = make([]responseOutput, len(ins.Requests))
	)
	if ins.Async {
		var wg sync.WaitGroup
		wg.Add(len(ins.Requests))
		for i, input := range ins.Requests {
			go func(i int, input requestsInput, wg *sync.WaitGroup) {
				defer wg.Done()
				var resp = s.request(input, taskData.ExecutionID)
				responses[i] = resp
			}(i, input, &wg)
		}
		wg.Wait()
	} else {
		for i, input := range ins.Requests {
			var resp = s.request(input, taskData.ExecutionID)
			responses[i] = resp
		}
	}

	bytOutput, _ := json.Marshal(responses)
	return "success", string(bytOutput)
}

func (s Task) request(ins requestsInput, exID string) responseOutput {
	req, err := http.NewRequest("POST", ins.URL, bytes.NewBuffer([]byte(ins.Body)))
	if err != nil {
		err = fmt.Errorf("error preparing HTTP POST for task with ExID [%v]: %v", exID, err)
		return responseOutput{Success: successOutput{0, ""}, Error: errorOutput{err.Error()}}
	}

	postRes, err := s.HTTPClient.Do(req)
	if err != nil {
		err = fmt.Errorf("error running HTTP POST for task with ExID [%v]: %v", exID, err)
		return responseOutput{Success: successOutput{0, ""}, Error: errorOutput{err.Error()}}
	}

	body, err := ioutil.ReadAll(postRes.Body)
	if err = postRes.Body.Close(); err != nil {
		err = fmt.Errorf("error closing response payload for task with ExID [%v]: %v", exID, err)
		return responseOutput{Success: successOutput{postRes.StatusCode, string(body)}, Error: errorOutput{err.Error()}}
	}

	return responseOutput{Success: successOutput{postRes.StatusCode, string(body)}, Error: errorOutput{""}}
}
