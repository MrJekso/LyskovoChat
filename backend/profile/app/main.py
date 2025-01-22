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
Профиль клиента
'''
@app.post('/profile')
async def profile(item: JWT):
    return {"response": item}
'''
Открыть профиль человека 
'''
@app.post('/profile/{item_id}')
async def open_profile(item_id: int,item: JWT):
    return {"response": item}
'''
Добавить в друзья
'''
@app.post('/profile/{item_id}/append')
async def add_profile(item_id: int,item: JWT):
    return {"response": item}
'''
Открыть чата с человеком
'''
@app.post('/profile/{item_id}/chat')
async def chat_profile(item_id: int,item: JWT):
    return {"response": item}
'''
Удалить друга
'''
@app.post('/profile/{item_id}/delete')
async def profile(item_id: int,item: JWT):
    return {"response": item}
'''
Поставить лайк
'''
@app.post('/profile/{item_id}/like')
async def set_like_profile(item_id: int,item: JWT):
    return {"response": item}
'''
Поставить дизлайк
'''
@app.post('/profile/{item_id}/dislike')
async def set_dislike_profile(item_id: int,item: JWT):
    return {"response": item}
