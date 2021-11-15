# Echo pair

Simple tool for debugging connection issues.

## Building

```
$ go build -o echopair main.go
```

## Running
### Server
The server is a simple TCP echo server

```
$ ./echopair serve --port 12345
```

### Client
This command repeatedly opens a connection to a echo server, sends a random 
string, reads the reply and compares is.

```
$ ./echopair client --address 127.0.0.1:12345 --interval 1s
```