from fastapi import FastAPI
from pydantic import BaseModel
from typing import Union

def check_jwt(jwt):
    pass
def create_jwt():
    pass

app = FastAPI()

class JWT(BaseModel):
    device: str
    jwt: str

@app.get("/healthcheck")
def healthcheck():
    return {"status": "ok"}
'''
Авторизация
'''
@app.post('/authorization')
async def authorization(item: JWT):
    return {"response": item}