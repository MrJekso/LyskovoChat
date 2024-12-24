from fastapi import FastAPI

app = FastAPI()

@app.get("/init")
def init():
    return {"message": "create db"}

@app.get("/create")
def create():
    return {"message": "generate data"}
