from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def home():
    return {"message": "home"}

@app.get("/test")
def test():
    return {"message": "test"}

@app.get("/frineds")
def test():
    return {"message": "frineds"}
