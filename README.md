# REST API

## Enpoints:

```
  GET  /ping  -> response: "pong"
  POST /arith -> response:  status code
```

This REST API receive requests containing tasks from producers, calls the appropriated subscriber to solve the problem and send back a response with the result.

## How to:

```
make test  // To run unit tests
make run   // To start the server
```

Server listens on port `8080`
