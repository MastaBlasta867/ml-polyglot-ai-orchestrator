import redis
import json
from fastapi import FastAPI
from sentence_transformers import SentenceTransformer

app = FastAPI()
r = redis.Redis(host="localhost", port=6379, db=0)

# Load embedding model once at startup
model = SentenceTransformer("all-MiniLM-L6-v2")

@app.get("/healthz")
def health_check():
    return {"status": "ok"}

@app.post("/process")
def process_document():
    # This will be replaced with real ML logic later
    return {"result": "Python processed the document successfully"}

def process_document_job(text: str):
    # Generate real embeddings
    embedding = model.encode(text).tolist()
    return {"embedding": embedding}

def check_for_jobs():
    job_data = r.lpop("ml_jobs")   # renamed
    if not job_data:
        return

    print("Job received from Redis")

    # Parse JSON job
    job = json.loads(job_data.decode())

    task = job.get("task")
    text = job.get("text")

    if task == "embed":
        result = process_document_job(text)
        r.rpush("ml_results", json.dumps(result))  # renamed

if __name__ == "__main__":
    while True:
        check_for_jobs()
