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
Список друзей
'''
@app.post('/friends')
async def list_friends(item: JWT):
    return {"response": item}
'''
Найти друга
'''
@app.post('/friends/find_friend')
async def find_friend(item: JWT):
    return {"response": item}
