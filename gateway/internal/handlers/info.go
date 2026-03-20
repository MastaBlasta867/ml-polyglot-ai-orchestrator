package handlers

import (
    "encoding/json"
    "net/http"
)

type InfoResponse struct {
    Service string `json:"service"`
    Status  string `json:"status"`
}

func Info(w http.ResponseWriter, r *http.Request) {
    response := InfoResponse{
        Service: "gateway",
        Status:  "ok",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

