FROM ubuntu:22.04
WORKDIR /app
COPY . .

RUN "pip install -r requirements.txt"

RUN "ls"
CMD ["/bin/ls"]
