# Greeting Example
This example shows a basic Go function named `GreetSomeone`, included in a Go package named `app`.
It takes a string, representing a person's name, as input. It returns a greeting that includes
this name as output.

The example also includes a file `start/main.go` that will invoke this function, using a name
supplied as a command-line argument. For example, you may run this program by passing the name
"Donna" to the program:

```
$ go run start/main.go Donna
```

This will output the greeting:

```
Hello Donna!
```

Although this example is not a Temporal Workflow, you will use this code
to create a Temporal Workflow during the first exercise in this course.
