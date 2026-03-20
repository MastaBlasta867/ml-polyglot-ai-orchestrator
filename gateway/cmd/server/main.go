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
    http.HandleFunc("/api/v1/info", handlers.Info)
    http.HandleFunc("/api/v1/documents", handlers.CreateDocumentJob)
    http.HandleFunc("/api/v1/jobs", handlers.GetJob)

    // Open the restaurant
    fmt.Println("Gateway listening on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
