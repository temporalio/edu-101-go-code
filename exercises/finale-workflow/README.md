# Exercise 4: Finale Workflow
One of the benefits of Temporal is that it provides SDKs for several
languages and you can multiple languages in the context of a single
Workflow. For example, you might write your Workflow in Go but use
Java or TypeScript for an Activity in that workflow. 

The last exercise in this workshop does exactly that. The Workflow
itself is written in Go, but the Activity that is executed as part
of this Workflow is written in Java, as is the Worker that runs it.
Since the Activity is written in Java, it's able to use a Java graphics
library that would otherwise be would be incompatible with a typical 
Go program.




# Start the Worker (Java)
In one terminal, run the following command:

```
$ java -classpath \
   java-activity-and-worker.jar \
   io.temporal.training.PdfCertWorker
```

# Start the Workflow (Go)
In another terminal, run the following command,
but be sure to replace `Sally Jiang` with your
full name.

```
$ go run start/main.go "Sally Jiang"
```

* Once the workflow is complete, use the explorer
view on the left side of the exercise environment
to locate the file created by this workflow. It
will have a name similar to `101_certificate_sally_jiang.pdf`,
only with your name in place of `sally_jiang`.
* Right-click its icon in the explorer view and choose
**Download...**. 
* After you've downloading it to your
computer, open it with your preferred PDF viewer.

Thank you for participating in Temporal 101!

### This is the end of the exercise.
