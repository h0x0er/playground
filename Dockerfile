FROM ubuntu:14.04
WORKDIR /app
COPY . .
RUN "ls"
CMD ["/bin/ls"]
