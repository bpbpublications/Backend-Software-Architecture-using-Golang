package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
)

func main() {
    srv := &http.Server{Addr: ":8080"}

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    // Wait for an interrupt signal
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)

    <-stop // Block until signal is received

    // Waiting for random time to shutdown in a controlled test.
    // Note: One can take up as an exercise adding a graceful
    // shutdown which waits for in progress tasks to complete.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    log.Println("Server exiting")
}
