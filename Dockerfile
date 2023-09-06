# our base image
FROM alpine:3.5

# Install python and pip
RUN apk add --update py2-pip

RUN pip install requests


# run the application
CMD ["/bin/ls"]