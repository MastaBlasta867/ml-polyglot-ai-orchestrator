package main

import (
    "fmt"
    "net/http"
    "github.com/MastaBlasta867/polyglot-ai-orchestrator/gateway/internal/handlers"
)

func main() {
    // Healthz worker
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })
    
    // Ping worker (placed AFTER, not INSIDE)
    http.HandleFunc("/ping", handlers.Ping)
    
    // Open the restaurant
    fmt.Println("Gateway listening on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
