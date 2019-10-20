# socklog

A basic websocket logging util library to help debug terminalUI applications.

Built upon [Gorilla's Websocket package.]("github.com/gorilla/websocket")

## Installation

```bash
go install "github.com/tomatosource/socklog"
```

## Usage

Start the server:

```bash
$ socklog
Listening for logs...

```

Initialise your logger:
```go
socklogger := socklog.MustNewLogger("localhost:8080")
defer socklogger.Close()
log.SetOutput(socklogger)
```

Log as normal:

```go
log.Print("hello world")
```
