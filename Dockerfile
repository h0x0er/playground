# our base image
FROM alpine:3.5

# Install python and pip
RUN apk add --update py2-pip

WORKDIR /test

COPY . .

RUN pip install -r requirements.txt


# run the application
CMD ["/bin/ls"]