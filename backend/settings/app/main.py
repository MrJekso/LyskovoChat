from fastapi import FastAPI
from pydantic import BaseModel
from typing import Union

def check_jwt(jwt):
    pass
def create_jwt():
    pass

app = FastAPI()

@app.get("/healthcheck")
def healthcheck():
    return {"status": "ok"}

class JWT(BaseModel):
    device: str
    jwt: str

@app.get('/settings')
def settings():
    return {"status":"ok","page":"settings"}
