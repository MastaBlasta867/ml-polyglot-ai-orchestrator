package handlers

import (
    "bytes"
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

type JobStatus string

const (
    JobStatusPending   JobStatus = "pending"
    JobStatusCompleted JobStatus = "completed"
)

type Job struct {
    ID     string    `json:"id"`
    Status JobStatus `json:"status"`
    Result string    `json:"result,omitempty"`
}

var (
    jobs   = make(map[string]*Job)
    jobsMu sync.RWMutex
)

type createDocumentRequest struct {
    Document string `json:"document"`
}

type createDocumentResponse struct {
    JobID string `json:"job_id"`
}

func newJobID() string {
    rand.Seed(time.Now().UnixNano())
    return fmt.Sprintf("job-%d", rand.Int63())
}

// POST /api/v1/documents
func CreateDocumentJob(w http.ResponseWriter, r *http.Request) {
    var req createDocumentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    id := newJobID()

    // 1. Create job in "pending" state
    job := &Job{
        ID:     id,
        Status: JobStatusPending,
    }

    jobsMu.Lock()
    jobs[id] = job
    jobsMu.Unlock()

    // 2. Kick off async processing
    go func(jobID string, document string) {
        // Prepare payload for Python service
        payload := map[string]string{"document": document}
        jsonData, _ := json.Marshal(payload)

        // Call Python service
        resp, err := http.Post("http://localhost:8000/process", "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            jobsMu.Lock()
            jobs[jobID].Status = JobStatusCompleted
            jobs[jobID].Result = fmt.Sprintf("error calling python service: %v", err)
            jobsMu.Unlock()
            return
        }
        defer resp.Body.Close()

        var pyResp struct {
            Result string `json:"result"`
        }
        json.NewDecoder(resp.Body).Decode(&pyResp)

        // 3. Update job with result
        jobsMu.Lock()
        jobs[jobID].Status = JobStatusCompleted
        jobs[jobID].Result = pyResp.Result
        jobsMu.Unlock()

    }(id, req.Document)

    // 4. Return job ID immediately
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createDocumentResponse{JobID: id})
}

// GET /api/v1/jobs?id=job-123
func GetJob(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "missing id", http.StatusBadRequest)
        return
    }

    jobsMu.RLock()
    job, ok := jobs[id]
    jobsMu.RUnlock()

    if !ok {
        http.Error(w, "job not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(job)
}
