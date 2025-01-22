from fastapi import FastAPI
from pydantic import BaseModel
from typing import Union

def check_jwt(jwt):
    pass
def create_jwt():
    pass

app = FastAPI()

################################
@app.get("/healthcheck")
def healthcheck():
    return {"status": "ok"}


################################
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
################################
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
################################
class JWT(BaseModel):
    device: str
    jwt: str
################################
'''
Профиль клиента
'''
@app.post('/profile')
async def profile(item: JWT):
    return {"response": item}
################################
'''
Список друзей
'''
@app.post('/friends')
async def list_friends(item: JWT):
    return {"response": item}
'''
Найти друга
'''
@app.post('/find_friend')
async def find_friend(item: JWT):
    return {"response": item}
################################
class Open_profile_people(BaseModel):
    item_id: str
    jwt: str
    device: str
################################
'''
Открыть профиль человека 
'''
@app.post('/people/{item_id}')
async def open_profile_people(item_id: int,item: JWT):
    return {"response": item}
################################
'''
Добавить в друзья
'''
@app.post('/people/{item_id}/append')
async def add_people(item_id: int,item: JWT):
    return {"response": item}
################################
'''
Открыть чата с человеком
'''
@app.post('/people/{item_id}/chat')
async def chat_people(item_id: int,item: JWT):
    return {"response": item}
################################
'''
Удалить друга
'''
@app.post('/people/{item_id}/delete')
async def people(item_id: int,item: JWT):
    return {"response": item}
################################
'''
Поставить лайк
'''
@app.post('/people/{item_id}/like')
async def set_like_people(item_id: int,item: JWT):
    return {"response": item}
################################
'''
Поставить дизлайк
'''
@app.post('/people/{item_id}/dislike')
async def set_dislike_people(item_id: int,item: JWT):
    return {"response": item}

'''
Список чатов и групп
'''
@app.post('/chats')
async def list_chat(item: JWT):
    return {"response": item}
################################
'''
Открыть чат
'''
@app.post('/chats/{chats_id}')
async def chat(chats_id: int, item: JWT):
    return {"response": item}
################################
'''
Написать сообщение
'''
@app.post('/chats/{chats_id}/send')
async def send_message(chats_id: int, item: JWT):
    return {"response": item}
################################
'''
Удалить чат
'''
@app.post('/chats/{chats_id}/delete')
async def delete_chat(chats_id: int, item: JWT):
    return {"response": item}
################################
'''
Добавить в чат
'''
@app.post('/chats/{chats_id}/append')
async def append_in_chat(chats_id: int, item: JWT):
    return {"response": item}
################################
'''
Изменить название группы/чаты
'''
@app.post('/chats/{chats_id}/edit_name')
async def edit_name_chat(chats_id: int, item: JWT):
    return {"response": item}
################################
'''
Изменить фото группы/чаты 
'''
@app.post('/chats/{chats_id}/edit_photo')
async def edit_photo(chats_id: int, item: JWT):
    return {"response": item}
################################
'''
Очистить группу/чат 
'''
@app.post('/chats/{chats_id}/clear')
async def clear_chat(chats_id: int, item: JWT):
    return {"response": item}
################################
'''
Уведомление от чатов Вкл\Выкл
'''
@app.post('/chats/{chats_id}/notifications')
async def set_notifications(chats_id: int, item: JWT):
    return {"response": item}
################################
@app.get('/settings')
def settings():
    return {"status":"ok","page":"settings"}
