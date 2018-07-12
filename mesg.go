package main

import (
	"context"
	"log"

	api "github.com/mesg-foundation/core/api/service"
	"google.golang.org/grpc"
)

// Fail early if can't connect
func mustNewServiceClient(mesgEndpoint string) api.ServiceClient {
	connection, err := grpc.Dial(mesgEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("mesg: couldn't dial grpc: %v", err)
	}
	return api.NewServiceClient(connection)
}

// Fail early if can't listen
func mustListenTaskRequests(cli api.ServiceClient, mesgToken string) api.Service_ListenTaskClient {
	var stream, err = cli.ListenTask(context.Background(), &api.ListenTaskRequest{
		Token: mesgToken,
	})
	if err != nil {
		log.Fatalf("mesg: couldn't listen for tasks: %v", err)
	}
	return stream
}
