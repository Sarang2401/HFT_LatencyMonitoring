from fastapi import FastAPI
from prometheus_client import Counter, generate_latest

app = FastAPI()

blocked_ticks = Counter(
    "blocked_ticks_total",
    "Number of ticks blocked by risk guard"
)

@app.post("/check")
def check_tick(price: float):
    if price > 1_000_000:
        blocked_ticks.inc()
        return {"status": "blocked"}
    return {"status": "ok"}

@app.get("/metrics")
def metrics():
    return generate_latest()
