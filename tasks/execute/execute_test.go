package execute

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	api "github.com/mesg-foundation/core/api/service"
)

func TestExecute(t *testing.T) {
	var ts = []struct {
		testName       string
		taskData       api.TaskData
		expectedKey    string
		expectedOutput string
	}{
		{
			testName: "happy case",
			taskData: api.TaskData{
				ExecutionID: "9DF2053F-3764-496A-9280-55EADDE6F52F",
				TaskKey:     "execute",
				InputData:   `{"url":"__URL__","body":"sample request"}`,
			},
			expectedKey:    "success",
			expectedOutput: `{"statusCode":200,"body":"sample response"}`,
		},
		{
			testName: "invalid url",
			taskData: api.TaskData{
				ExecutionID: "9DF2053F-3764-496A-9280-55EADDE6F52F",
				TaskKey:     "execute",
				InputData:   `{"url":"invalid_url","body":"sample request"}`,
			},
			expectedKey:    "error",
			expectedOutput: `{"message":"error running HTTP POST for task with ExID [9DF2053F-3764-496A-9280-55EADDE6F52F]: Post invalid_url: unsupported protocol scheme \"\""}`,
		},
		{
			testName: "invalid input data",
			taskData: api.TaskData{
				ExecutionID: "9DF2053F-3764-496A-9280-55EADDE6F52F",
				TaskKey:     "execute",
				InputData:   `not a json`,
			},
			expectedKey:    "error",
			expectedOutput: `{"message":"error unmarshalling input of task with ExID [9DF2053F-3764-496A-9280-55EADDE6F52F]: invalid character 'o' in literal null (expecting 'u')"}`,
		},
	}
	for _, tc := range ts {
		t.Run(tc.testName, func(t *testing.T) {
			var (
				server    = httptest.NewServer(testHandler{})
				processor = Task{HTTPClient: *server.Client()}
			)
			tc.taskData.InputData = strings.Replace(tc.taskData.InputData, "__URL__", server.URL, -1)
			var actualKey, actualOutput = processor.Process(&tc.taskData)

			if tc.expectedKey != actualKey {
				t.Errorf("expected key [%v] but got [%v]", tc.expectedKey, actualKey)
			}
			if tc.expectedOutput != actualOutput {
				t.Errorf("expected output [%v] but got [%v]", tc.expectedOutput, actualOutput)
			}
		})
	}
}

type testHandler struct{}

func (h testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "sample response")
}
