FROM golang:latest
RUN mkdir memo-app
WORKDIR /memo-app

RUN apt-get update \
  && apt-get install -y mariadb-client

COPY . .

RUN CGO_ENABLED=0 go build -o main src/web.go

CMD ["./main"]