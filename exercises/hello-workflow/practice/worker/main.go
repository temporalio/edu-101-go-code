package main

import (
	"log"
	hello "temporal101/exercises/hello-workflow/practice"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// TODO: modify the statement below to specify the task queue name
	w := worker.New(c, "greeting-tasks", worker.Options{})

	w.RegisterWorkflow(hello.GreetSomeone)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
