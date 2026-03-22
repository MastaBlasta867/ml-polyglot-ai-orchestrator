from fastapi import FastAPI

app = FastAPI()

@app.get("/healthz")
def health_check():
    return {"status": "ok"}

@app.post("/process")
def process_document():
    return {"message": "Python service received the request"}