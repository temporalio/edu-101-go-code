# Exercise 3: Farewell Workflow
During this exercise, you will create an Activity function,
register it with the Worker, and modify a Workflow to execute it.

Before continuing, kill any Worker instances still running from any
previous exercises (you can do this by pressing Ctrl-C in the terminals 
where they are running.

As with other exercises, you should make your changes in the `practice` 
directory. Look for "TODO" comments, which will provide hints about what
you'll need to change. If you get stuck and need additional hints, or 
if you want to check your work, look at the completed example in the
`solution` directory. 

## Part A: Write an Activity Function
The `translate.go` file contains a function (`GreetInSpanish`) that uses a 
microservice to get a customized greeting in Spanish. This file also contains 
a utility function ('callService') that the Activity uses to call the microservice. 

1. Open the `translate.go` file (located in the `practice` subdirectory) in the editor
2. Create a new Activity that will get a custom farewell message from the microservice.
   1. Copy the `GreetInSpanish` function
   2. Rename the new function (use any valid name you like)
   3. Change '`get-spanish-greeting`" in this new function to "`get-spanish-farewell`"
   4. Save your changes to this file

## Part B: Register the Activity Function
1. Open the `main.go` file (located in the `practice/worker` subdirectory) in the editor
2. Add a new line that registers your new Activity with the Worker (hint: you'll use the
   name you used for the new function in an earlier step).
3. Save your changes to this file


## Part C: Modify the Workflow to Execute Your New Activity
1. Open the `greeting.go` file (located in the `practice` subdirectory) in the editor
2. Look for the TODO message, uncomment the line below it, and then change that line
   to execute the Activity function you created
3. Save your changes to this file


## Part D: Start the Microservice and Run the Workflow
All commands below must be run from the `practice` subdirectory.

1. Start the microservice by running `go run microservice/greeting-and-farewell-service.go` in a terminal
2. In another terminal, start your Worker by running `go run worker/main.go`
3. In a third terminal, execute your Workflow by running `go run start/main.go Donna` 
    (replacing `Donna` with your own name)

If there is time remaining, experiment with Activity failures and retries 
by stopping the microservice (press Ctrl-C in its terminal) and re-running 
the Workflow. Look at the Web UI to see the status of the Workflow and its
Activities. After a few seconds, restart the microservice by running the
same command used to start it earlier. You should find that the Workflow
will now complete successfully following the next Activity retry.
