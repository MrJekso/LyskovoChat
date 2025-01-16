from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

################################
@app.get("/healthcheck")
def healthcheck():
    return {"status": "ok"}
################################
@app.get('/profile')
def profile():
    return {"status":"ok","page":"profile"}
@app.get('/friends')
def friends():
    return {"status":"ok","page":"friends"}
@app.get('/news')
def news():
    return {"status":"ok","page":"news"}
@app.get('/settings')
def settings():
    return {"status":"ok","page":"settings"}
@app.get('/message')
def message():
    return {"status":"ok","page":"message"}

class Item(BaseModel):
    email: str
    login: str
    password: str
    repeat_password: str

@app.post('/registration')
async def registration(item: Item):
    return {"response": "Item.email"}