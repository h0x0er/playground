FROM python:3.11
WORKDIR /app
COPY . .

RUN "ls"

RUN "python -m pip install -r requirements.txt"

CMD ["/bin/ls"]
