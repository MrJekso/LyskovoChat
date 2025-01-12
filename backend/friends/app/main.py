from fastapi import FastAPI

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

