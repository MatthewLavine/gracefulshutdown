# gracefulshutdown

[![Go](https://github.com/MatthewLavine/gracefulshutdown/actions/workflows/go.yml/badge.svg)](https://github.com/MatthewLavine/gracefulshutdown/actions/workflows/go.yml)

`gracefulshutdown` is a Go library that allows applications to register shutdown handlers to be executed during server shutdown.

## Usage

Register a shutdown handler:

```go
gracefulshutdown.AddShutdownHandler(func() error {
    log.Println("Shutting down server")
    return httpServer.Shutdown(ctx)
})
```

See `example/example.go` for a full working example.
