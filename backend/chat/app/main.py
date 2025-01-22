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
Список чатов и групп
'''
@app.post('/chats')
async def list_chat(item: JWT):
    return {"response": item}
'''
Открыть чат
'''
@app.post('/chats/{chats_id}')
async def chat(chats_id: int, item: JWT):
    return {"response": item}
'''
Написать сообщение
'''
@app.post('/chats/{chats_id}/send')
async def send_message(chats_id: int, item: JWT):
    return {"response": item}
'''
Удалить чат
'''
@app.post('/chats/{chats_id}/delete')
async def delete_chat(chats_id: int, item: JWT):
    return {"response": item}
'''
Добавить в чат
'''
@app.post('/chats/{chats_id}/append')
async def append_in_chat(chats_id: int, item: JWT):
    return {"response": item}
'''
Изменить название группы/чаты
'''
@app.post('/chats/{chats_id}/edit_name')
async def edit_name_chat(chats_id: int, item: JWT):
    return {"response": item}
'''
Изменить фото группы/чаты 
'''
@app.post('/chats/{chats_id}/edit_photo')
async def edit_photo(chats_id: int, item: JWT):
    return {"response": item}
'''
Очистить группу/чат 
'''
@app.post('/chats/{chats_id}/clear')
async def clear_chat(chats_id: int, item: JWT):
    return {"response": item}
'''
Уведомление от чатов Вкл\Выкл
'''
@app.post('/chats/{chats_id}/notifications')
async def set_notifications(chats_id: int, item: JWT):
    return {"response": item}