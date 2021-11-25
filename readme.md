# KeyGo

KeyGo project is made to organize key-value managements in a simple way. Keep keys in memory to the until specified date.

## Installation

You can clone from the source;
```bash
$ git clone https://github.com/hakantaskin/keygo
```
Run & Build:
```bash
$ go run main.go
```
```bash
$ go build main.go
```

## Usage:
Get:
```bash
$ curl http://localhost:8585/get\?key\=deneme
```
```json
{"error":"record not found! key: deneme","data":null}
```

Set:
```bash
$ curl -i -XPOST  http://localhost:8585/set\?key\=deneme -d '{"xxxYzz": 5}'
```
```json
{"error":"","data":"{\"xxxYzz\": 5}"}
```

Del:
```bash
$ curl -i -XDELETE  http://localhost:8585/del\?key\=deneme
```
```json
{"error":"","data":"ok"}
```

Flush:
```bash
$ curl -i -XDELETE  http://localhost:8585/flush
```
```json
{"error":"","data":"ok"}
```

---

## Docker

build:

```bash
$ docker build . -t keygo
```

run:

```bash
$ docker run -dp 8585:8585 -i -t keygo
```

other commands: 

```bash
$ docker image ls
$ docker container ls
$ docker exec -it <container-id> /bin/sh
$ docker logs <container-id>
$ docker stop <container-id>
```

# Tests
You can run tests using
```bash
$ go test ./...
$ go test -cover ./...
```