package main

import (
	"context"
	"log"

	api "github.com/mesg-foundation/core/api/service"
)

type task interface {
	Process(taskData *api.TaskData) (string, string)
}

type streamProcessor struct {
	cli    api.ServiceClient
	stream api.Service_ListenTaskClient
	tasks  map[string]task
}

func newStreamProcessor(cli api.ServiceClient, stream api.Service_ListenTaskClient, tasks map[string]task) streamProcessor {
	return streamProcessor{cli, stream, tasks}
}

func (s streamProcessor) mustStart() {
	for {
		res, err := s.stream.Recv()
		if err != nil {
			log.Fatalf("stream: error receiving tasks (or stream closed; can't continue): %v\n", err)
		}

		task, ok := s.tasks[res.TaskKey]
		if !ok {
			log.Printf("stream: didn't find task with key %v\n", res.TaskKey)
			continue
		}

		var outputKey, outputData = task.Process(res)

		reply, err := s.cli.SubmitResult(context.Background(), &api.SubmitResultRequest{
			ExecutionID: res.ExecutionID,
			OutputKey:   outputKey,
			OutputData:  outputData,
		})
		if err != nil {
			log.Printf("stream: error submitting result of task with ExID [%v]: %v; ignoring.\n", res.ExecutionID, err)
		}
		log.Println(reply)
	}
}
