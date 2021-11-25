FROM golang:latest

RUN mkdir /app
WORKDIR /app
ADD . /app 

RUN cd /app &&  go build
CMD ./docker-training

