from fastapi import FastAPI
import requests
import json


with open('config.json') as f:
    templates = json.load(f)
ip,port   = templates.items()
ip,port = ip[1],port[1]

app = FastAPI()
#######################################
@app.get("/")
def healthcheck():
    return {"status":"ok"}
#######################################
@app.get("/healthcheck")
def healthcheck():
    url = "http://" + ip + ":" + port + "/healthcheck"
    r = requests.get(url)
    return {"response": r.text}
#######################################
@app.get('/profile')
def profile():
    url = "http://" + ip + ":" + port + "/profile"
    r = requests.get(url)
    return {"response": r.text}

@app.get('/friends')
def friends():
    url = "http://" + ip + ":" + port + "/friends"
    r = requests.get(url)
    return {"response": r.text}

@app.get('/news')
def news():
    url = "http://" + ip + ":" + port + "/news"
    r = requests.get(url)
    return {"response": r.text}

@app.get('/settings')
def settings():
    url = "http://" + ip + ":" + port + "/settings"
    r = requests.get(url)
    return {"response": r.text}

@app.get('/message')
def message():
    url = "http://" + ip + ":" + port + "/message"
    r = requests.get(url)
    return {"response": r.text}
