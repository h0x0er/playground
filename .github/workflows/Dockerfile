FROM ubuntu:14.04
WORKDIR /app
COPY . .
RUN "apt-get update && sudo apt-get install -y   wget   curl   git   dos2unix   software-properties-common"
CMD ["/bin/ls"]
