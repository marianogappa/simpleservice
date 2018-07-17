package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/marianogappa/simpleservice/tasks/execute" // Implemented tasks
	"github.com/marianogappa/simpleservice/tasks/execute_many"

	api "github.com/mesg-foundation/core/api/service"
	"github.com/satori/go.uuid"
)

func main() {
	var (
		cli        = mustNewServiceClient(os.Getenv("MESG_ENDPOINT"))
		stream     = mustListenTaskRequests(cli, os.Getenv("MESG_TOKEN"))
		httpClient = http.Client{Timeout: 30 * time.Second}
		tasks      = map[string]task{
			"execute":     execute.Task{HTTPClient: httpClient},
			"executeMany": execute_many.Task{HTTPClient: httpClient}, // Adding tasks: add entry here and task subpackage
		}
		streamProcessor = newStreamProcessor(cli, stream, tasks)
		emitEvent       = func(body string) (*api.EmitEventReply, error) {
			return cli.EmitEvent(context.Background(), &api.EmitEventRequest{
				Token:     os.Getenv("MESG_TOKEN"),
				EventKey:  "onRequest",
				EventData: body,
			})
		}
	)
	go streamProcessor.mustStart()
	mustServe(":8080", simpleEndpoint{now: now, uuid: mustNewUUID, emitEvent: emitEvent})
}

func now() string {
	return time.Now().UTC().Format("2006-01-02 03:04:05")
}

func mustNewUUID() string {
	var u, err = uuid.NewV1() // Guaranteed lexicographically sortable
	if err != nil {
		log.Fatalf("mustNewUUID: system couldn't generate UUID v1: %v", err)
	}
	return u.String()
}
