FROM golang:latest
MAINTAINER kshamiev konstantin@shamiev.ru

WORKDIR /usr/src/app
COPY . .

RUN go build -i -mod vendor -o bin/app .;
EXPOSE 8080

CMD bin/app -c config.yaml;
