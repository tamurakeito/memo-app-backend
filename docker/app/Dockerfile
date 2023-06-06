FROM golang:latest
RUN mkdir test-docker
WORKDIR /test-docker

RUN apt-get update \
  && apt-get install -y mariadb-client

COPY . .

RUN CGO_ENABLED=0 go build -o main main.go

CMD ["./main"]