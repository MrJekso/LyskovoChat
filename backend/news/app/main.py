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
"""
drop comment
"""
@app.get("/healthcheck")
def healthcheck():
    return {"status": "ok"}
'''
Список новостей
'''
@app.post('/news')
async def list_news(item: JWT):
    return {"response": item}
'''
Создание новости
'''
@app.post('/create_news')
async def create_news(item: JWT):
    return {"response": item}
'''
Список новостей
'''
@app.post('/news')
async def list_news(item: JWT):
    return {"response": item}
'''
Открыть комментарии
'''
@app.post('/news/{news_id}/comments')
async def open_comment_news(news_id: int, item: JWT):
    return {"response": item}
'''
Добавить комментарий
'''
@app.post('/news/{news_id}/comments/append')
async def append_comment_news(news_id: int, item: JWT):
    return {"response": item}
'''
Поставить лайки
'''
@app.post('/news/{news_id}/like')
async def set_dislaik_news(news_id: int, item: JWT):
    return {"response": item}
'''
Поставить дизлайк
'''
@app.post('/news/{news_id}/dislike')
async def set_dislaike_news(news_id: int, item: JWT):
    return {"response": item}
'''
Поставить жалобу
'''
@app.post('/news/{news_id}/complaint')
async def set_complaint(news_id: int, item: JWT):
    return {"response": item}