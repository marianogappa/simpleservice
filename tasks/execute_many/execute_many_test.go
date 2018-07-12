package execute_many

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	api "github.com/mesg-foundation/core/api/service"
)

func TestExecuteMany(t *testing.T) {
	var ts = []struct {
		testName       string
		taskData       api.TaskData
		expectedKey    string
		expectedOutput string
	}{
		{
			testName: "happy case: sync",
			taskData: api.TaskData{
				ExecutionID: "9DF2053F-3764-496A-9280-55EADDE6F52F",
				TaskKey:     "execute",
				InputData:   `{"requests":[{"url":"__URL__?v=1","body":"sample request"},{"url":"__URL__?v=2","body":"sample request"}],"async":false}`,
			},
			expectedKey:    "success",
			expectedOutput: `[{"success":{"statusCode":200,"body":"sample response 1"},"error":{"message":""}},{"success":{"statusCode":200,"body":"sample response 2"},"error":{"message":""}}]`,
		},
		{
			testName: "happy case: async",
			taskData: api.TaskData{
				ExecutionID: "9DF2053F-3764-496A-9280-55EADDE6F52F",
				TaskKey:     "execute",
				InputData:   `{"requests":[{"url":"__URL__?v=3","body":"sample request1"},{"url":"__URL__?v=4","body":"sample request2"}],"async":true}`,
			},
			expectedKey:    "success",
			expectedOutput: `[{"success":{"statusCode":200,"body":"sample response 3"},"error":{"message":""}},{"success":{"statusCode":200,"body":"sample response 4"},"error":{"message":""}}]`,
		},
		{
			testName: "one invalid url still causes success, with one error response",
			taskData: api.TaskData{
				ExecutionID: "9DF2053F-3764-496A-9280-55EADDE6F52F",
				TaskKey:     "execute",
				InputData:   `{"requests":[{"url":"invalid url","body":"sample request1"},{"url":"__URL__?v=6","body":"sample request2"}],"async":true}`,
			},
			expectedKey:    "success",
			expectedOutput: `[{"success":{"statusCode":0,"body":""},"error":{"message":"error running HTTP POST for task with ExID [9DF2053F-3764-496A-9280-55EADDE6F52F]: Post invalid%20url: unsupported protocol scheme \"\""}},{"success":{"statusCode":200,"body":"sample response 6"},"error":{"message":""}}]`,
		},
		{
			testName: "invalid json causes isSuccess: false with main error",
			taskData: api.TaskData{
				ExecutionID: "9DF2053F-3764-496A-9280-55EADDE6F52F",
				TaskKey:     "execute",
				InputData:   `invalid json`,
			},
			expectedKey:    "error",
			expectedOutput: `{"message":"error unmarshalling input of task with ExID [9DF2053F-3764-496A-9280-55EADDE6F52F]: invalid character 'i' looking for beginning of value"}`,
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
	io.WriteString(w, fmt.Sprintf("sample response %v", r.URL.Query().Get("v")))
}
