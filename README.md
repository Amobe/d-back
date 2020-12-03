D-Back
---

D-Back is a practice project, aim to provide the request limited web service.

## Description

The d-back server has the request rate limiter which count the number of request of each duration.
The server can be configured with environment variables PORT, NUMBER, and DURATION.
For example, set NUMBER=10 and DURATION=1m means the server allow to accept 10 request in 1 minute.
If there is no request incoming, the limiter with be pause until new request coming to save the resource.

## Development environment

```
OS: macOS Big Sur (11.0.1)
Golang: 1.15.5
Docker: 2.5.0.1
```

## How to test

Execute go test with coverage.
```bash
make test
```

## How to start

### From source

Start the d-back server.
```bash
PORT=10080 NUMBER=60 DURATION=1m go run ./cmd/d-back/main.go
```

Start the d-back client.
```bash
# times: specific the request times, -1 means endless mode.
go run ./cmd/d-back-client/main.go -url=http://127.0.0.1:10080/hello -times=-1 -interval=1s
```

Sample output:
```
// when request is accepted
...
status: 200 OK, body: {"index":57}
status: 200 OK, body: {"index":58}
status: 200 OK, body: {"index":59}
status: 200 OK, body: {"index":60}
// when request times in duration is over the limit
status: 503 Service Unavailable, body: {"message":"request token: accpectance denied"}
status: 503 Service Unavailable, body: {"message":"request token: accpectance denied"}
```

### From docker

Build the image.
```bash
make docker.build
```

Start the d-back server.
```bash
docker run -it --rm -p 10080:10080 --name d-back d-back
```

Start the d-back client.
```bash
docker run -it --rm --link d-back d-back-client -url="http://d-back:10080/hello" -times=-1 -interval=1s
```

### Complex test

#### Racing request with same IP

```
Give 1 docker server and 2 source clients
When clients request at same times
Then clients should racing the request number
```

#### Request number distinguishes by different IP

```
Give 1 docker server and 2 docker clients
When clients request at same times
Then clients should not racing the request number
```
