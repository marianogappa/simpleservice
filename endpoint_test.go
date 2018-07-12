package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	api "github.com/mesg-foundation/core/api/service"
)

func TestEndpoint(t *testing.T) {
	var ts = []struct {
		testName           string
		now                string
		uuid               string
		requestBody        string
		emitError          error
		httpMethod         string
		expectedStatusCode int
		expectedReply      string
	}{
		{
			testName:           "GET method is not allowed",
			now:                "2018-07-12 01:02:03",
			uuid:               "BEFB796D-4C63-46FD-9AB6-84D1307DEE5F",
			requestBody:        "Random request payload to relay",
			emitError:          nil,
			httpMethod:         "GET",
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedReply:      ``,
		},
		{
			testName:           "an emit error causes an InternalServerError result, but could still relay",
			now:                "2018-07-12 01:02:03",
			uuid:               "BEFB796D-4C63-46FD-9AB6-84D1307DEE5F",
			requestBody:        "Random request payload to relay",
			emitError:          fmt.Errorf("for test"),
			httpMethod:         "POST",
			expectedStatusCode: http.StatusInternalServerError,
			expectedReply:      `{"date":"2018-07-12 01:02:03","id":"BEFB796D-4C63-46FD-9AB6-84D1307DEE5F","body":"Random request payload to relay"}`,
		},
		{
			testName:           "no emitted errors returns HTTP OK and proper replies",
			now:                "2018-07-12 01:02:03",
			uuid:               "BEFB796D-4C63-46FD-9AB6-84D1307DEE5F",
			requestBody:        "Random request payload to relay",
			emitError:          nil,
			httpMethod:         "POST",
			expectedStatusCode: http.StatusOK,
			expectedReply:      `{"date":"2018-07-12 01:02:03","id":"BEFB796D-4C63-46FD-9AB6-84D1307DEE5F","body":"Random request payload to relay"}`,
		},
		{
			testName:           "another happy case, but checking that a different date is emitted properly",
			now:                "2001-02-03 04:05:06",
			uuid:               "BEFB796D-4C63-46FD-9AB6-84D1307DEE5F",
			requestBody:        "Random request payload to relay",
			emitError:          nil,
			httpMethod:         "POST",
			expectedStatusCode: http.StatusOK,
			expectedReply:      `{"date":"2001-02-03 04:05:06","id":"BEFB796D-4C63-46FD-9AB6-84D1307DEE5F","body":"Random request payload to relay"}`,
		},
		{
			testName:           "another happy case, but checking that a different uuid is emitted properly",
			now:                "2018-07-12 01:02:03",
			uuid:               "0F48AAA7-EEFB-4699-BDD5-791B366D0294",
			requestBody:        "Random request payload to relay",
			emitError:          nil,
			httpMethod:         "POST",
			expectedStatusCode: http.StatusOK,
			expectedReply:      `{"date":"2018-07-12 01:02:03","id":"0F48AAA7-EEFB-4699-BDD5-791B366D0294","body":"Random request payload to relay"}`,
		},
		{
			testName:           "another happy case, but checking that a different payload is emitted properly",
			now:                "2018-07-12 01:02:03",
			uuid:               "BEFB796D-4C63-46FD-9AB6-84D1307DEE5F",
			requestBody:        "Different request payload to relay",
			emitError:          nil,
			httpMethod:         "POST",
			expectedStatusCode: http.StatusOK,
			expectedReply:      `{"date":"2018-07-12 01:02:03","id":"BEFB796D-4C63-46FD-9AB6-84D1307DEE5F","body":"Different request payload to relay"}`,
		},
	}

	for _, tc := range ts {
		t.Run(tc.testName, func(t *testing.T) {
			var (
				r           = httptest.NewRequest(tc.httpMethod, "http://example.com/foo", strings.NewReader(tc.requestBody))
				w           = httptest.NewRecorder()
				actualReply string
			)
			simpleEndpoint{
				now:  func() string { return tc.now },
				uuid: func() string { return tc.uuid },
				emitEvent: func(body string) (*api.EmitEventReply, error) {
					actualReply = body
					return nil, tc.emitError
				},
			}.ServeHTTP(w, r)

			var resp = w.Result()
			if tc.expectedStatusCode != resp.StatusCode {
				t.Errorf("expected statuscode %v but got %v", tc.expectedStatusCode, resp.StatusCode)
			}
			if tc.expectedReply != actualReply {
				t.Errorf("expected reply %v but got %v", tc.expectedReply, actualReply)
			}
		})
	}
}
