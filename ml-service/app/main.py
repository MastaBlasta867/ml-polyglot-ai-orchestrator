import redis

from fastapi import FastAPI

app = FastAPI()
r = redis.Redis(host="localhost", port=6379, db=0)

@app.get("/healthz")
def health_check():
    return {"status": "ok"}

@app.post("/process")
def process_document():
    # This is where real processing will go later
    return {"result": "Python processed the document successfully"}

def make_pizza(order):
    # This is where real ML work will go later
    return f"Pizza for order {order} is ready!"

def check_for_orders():
    order = r.lpop("pizza_orders")
    if order:
        result = make_pizza(order.decode())
        r.rpush("pizza_results", result)

if __name__ == "__main__":
    while True:
        check_for_orders()
