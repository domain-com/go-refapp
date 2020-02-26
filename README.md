# go-refapp

This is a sample application in Go.

## Building

To build the application, you must have Go already installed or build the
Dockerfile. We assume that you have your Go installation working.

### Compile with Go

You have to install the dependencies with go get.

```bash
go get -v ./...
```

Then, build the application.

```bash
go build -i -v
```

### Build the Dockerfile

You have to build the Dockerfile.

```bash
docker build -t go-refapp:0.0.1 .
```

## Running

### Run the binary

Execute the binary.

```bash
$ ./go-refapp
2020/02/10 12:10:19 Listening port: 8080
```

### Run the docker image

Execute docker run with the port 8080

```bash
$ docker run -it --rm --name go-refapp -p 8080:8080 go-refapp:0.0.1
2020/02/10 15:13:17 Listening port: 8080
```

## Access to the sample application

This application is just a sample. Provides the endpoint `articles`.
You can test with curl.

```bash
$ curl http://localhost:8080/articles
[{"Title":"Test Title","description":"Test Description","content":"Test Content"}]
```