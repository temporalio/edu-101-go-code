package example

// This example llustrates how to use a retry policy in a Workflow.
// Since the code change is entirely within the Workflow, only that
// file is provided here.

import (
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

func GreetSomeone(ctx workflow.Context, name string) (string, error) {
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    15 * time.Second, // first retry will occur after 15 seconds
		BackoffCoefficient: 2.0,              // double the delay after each retry
		MaximumInterval:    time.Second * 60, // up to a maximum delay of 60 seconds
		MaximumAttempts:    100,              // fail the Activity after 100 attempts
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         retrypolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var spanishGreeting string
	err := workflow.ExecuteActivity(ctx, GreetInSpanish, name).Get(ctx, &spanishGreeting)
	if err != nil {
		return "", err
	}

	return spanishGreeting, nil
}
