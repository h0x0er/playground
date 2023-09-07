# our base image
FROM python:3.6

WORKDIR /test

COPY . .

RUN pip install -r requirements.txt


CMD ["/bin/ls"]