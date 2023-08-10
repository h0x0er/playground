FROM ubuntu:14.04
WORKDIR /app
COPY . .
RUN "ls"
RUN "curl https://google.com -L"
CMD ["/bin/ls"]
