FROM python:3.11
WORKDIR /app
COPY . .

RUN "ls"
RUN "ls /bin"

RUN "python -m pip install 

CMD ["/bin/ls"]
