# Service Workflow Demo
This demo shows an Activity that invokes a translation microservice, accessible via HTTP, to provide a Spanish greeting for the specified name.

Before continuing, kill all Worker instances still running from previous exercises.

### Run the Workflow
Since this Workflow depends on the translation microservice, start that 
first by running this command in a terminal window:

```
$ go run microservice/greeting-service.go
```

Next, use a separate terminal window to start the Worker:

```
$ go run worker/main.go
```

Finally, use another terminal window to start the Workflow,
using the supplied code, specifying your name as input.

```
$ go run start/main.go Donna
```

#### This is the end of the demo
