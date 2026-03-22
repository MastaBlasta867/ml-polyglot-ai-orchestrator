package main

import (
    "fmt"
    "net/http"

    "github.com/MastaBlasta867/ml-polyglot-ai-orchestrator/gateway/internal/handlers"
    "github.com/MastaBlasta867/ml-polyglot-ai-orchestrator/gateway/internal/services"
)

func main() {
    // Create Python client
    pythonClient := services.NewPythonClient("http://localhost:8001")

    // Create document handler
    docHandler := handlers.NewDocumentHandler(pythonClient)

    // Existing handlers
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })
    http.HandleFunc("/ping", handlers.Ping)
    http.HandleFunc("/api/v1/info", handlers.Info)
    http.HandleFunc("/api/v1/documents", handlers.CreateDocumentJob)
    http.HandleFunc("/api/v1/jobs", handlers.GetJob)

    // NEW: Route that calls Python
    http.HandleFunc("/documents", docHandler.CreateDocument)

    // Start server
    fmt.Println("Gateway listening on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
