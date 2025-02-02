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
curl -X 'POST' \
  'http://url/registration' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "string",
  "login": "string",
  "password": "string",
  "repeat_password": "string"
}
'''
class Registration(BaseModel):
    email: str
    login: str
    password: str
    repeat_password: str
    device: str
@app.post('/registration')
async def registration(item: Registration):
    return {"response": item}
'''
Аунтификация
'''
class Authentication(BaseModel):
    login: str
    password: str
    device: str
@app.post('/authentication')
async def authentication(item: Authentication):
    return {"response": item}
