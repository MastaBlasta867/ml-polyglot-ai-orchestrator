package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type PythonClient struct {
    BaseURL string
}

func NewPythonClient(baseURL string) *PythonClient {
    return &PythonClient{BaseURL: baseURL}
}

func (c *PythonClient) ProcessDocument() (map[string]interface{}, error) {
    url := fmt.Sprintf("%s/process", c.BaseURL)

    // Empty JSON body for now
    body := bytes.NewBuffer([]byte("{}"))

    resp, err := http.Post(url, "application/json", body)
    if err != nil {
        return nil, fmt.Errorf("failed to call python service: %w", err)
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode python response: %w", err)
    }

    return result, nil
}
