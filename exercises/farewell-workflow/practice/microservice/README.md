microservice
============
This directory contains the Go code for a simple microservice, 
accessible via HTTP, which accepts a name as an input parameter 
and returns a greeting in Spanish customized for that name.

Run the service:

```
$ go run greeting-and-farewell-service.go
```

On a Mac, you might be prompted by the OS as to whether to allow incoming connections, so approve them if so.

Afterwards, test that it's working by visiting this URL in your browser:

[http://localhost:9999/get-spanish-greeting?name=Donna]

Once it is, you can write and run workflows that use this service.
