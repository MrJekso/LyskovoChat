FROM python:3 AS admin

WORKDIR /usr/src/app

COPY req.txt ./
RUN pip install --no-cache-dir -r req.txt
RUN pip install "fastapi[standard]"

COPY ./app/main.py ./main.py

CMD [ "fastapi", "run", "/usr/src/app/main.py" ]
