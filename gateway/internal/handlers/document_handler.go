package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/MastaBlasta867/ml-polyglot-ai-orchestrator/gateway/internal/services"
)

type DocumentHandler struct {
    Python *services.PythonClient
}

func NewDocumentHandler(pythonClient *services.PythonClient) *DocumentHandler {
    return &DocumentHandler{
        Python: pythonClient,
    }
}

func (h *DocumentHandler) CreateDocument(w http.ResponseWriter, r *http.Request) {
    result, err := h.Python.ProcessDocument()
    if err != nil {
        http.Error(w, "failed to call python service", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}
