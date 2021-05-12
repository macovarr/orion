# Golang Version of the Elevator Server
This doc should give you links to everything you need to know to develop
the elevator-server code in Go (currently tested in Go 1.16.4, Swagger v0.26.0).

# Generate
We are using the [`go-swagger`](github.com/go-swagger/go-swagger/cmd/swagger)
tool to build the generic scaffolding for our project.

First, let's get the repo.

```
$ export REPO_ROOT=<your_path_here>
$ git clone git@github.com:OnBeep/elevator-server.git $REPO_ROOT
$ cd $REPO_ROOT/go-ele
$ go get ./...
```
Further references to this project will be as
`$REPO_ROOT`.

Now, install the code generation tool:
```
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

The clone will come with a pre-generated API bundle. If you would like to re-generate
run the following command: 
```
$ cd $REPO_ROOT/go-ele
$ swagger generate server -f ../elevator.yml
```

Now run it:
```
$ go run $REPO_ROOT/cmd/lifty-server/main.go --port 5000
```

Alternatively install the server into your OS first
```
$ $GOPATH/bin/lifty-server --port 5000
>> serving lifty at http://127.0.0.1:5000
```

In another terminal or browser you should now see the server responding:
```
$ curl http://127.0.0.1:5000/v1/welcome
>> {"msg":"Welcome to the Elevator Server"}
```

To run the unit tests:
```
$ cd $REPO_ROOT/go-ele
$ go test -v ./...
```
