package farewell

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetSomeone(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var spanishGreeting string
	err := workflow.ExecuteActivity(ctx, GreetInSpanish, name).Get(ctx, &spanishGreeting)
	if err != nil {
		return "", err
	}

	var spanishFarewell string
	// TODO: uncomment the four lines below and change the first one to execute the 
	// Activity function you created
	//err = workflow.ExecuteActivity(ctx, GreetInSpanish, name).Get(ctx, &spanishFarewell)
	//if err != nil {
	//	return "", err
	//}

	var helloGoodbye = "\n" + spanishGreeting + "\n" + spanishFarewell

	return helloGoodbye, nil
}
