package main

import (
	"log"
	farewell "temporal101/exercises/farewell-workflow/solution"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "greeting-tasks", worker.Options{})

	w.RegisterWorkflow(farewell.GreetSomeone)
	w.RegisterActivity(farewell.GreetInSpanish)
	w.RegisterActivity(farewell.FarewellInSpanish)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
