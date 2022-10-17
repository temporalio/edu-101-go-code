package main

import (
	"context"
	"log"
	"os"
	finale "temporal101/exercises/finale-workflow"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "generate-certificate-workflow",
		TaskQueue: "generate-certificate-taskqueue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, finale.CertificateGeneratorWorkflow, os.Args[1])
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Generated certificate at: ", result)
}
