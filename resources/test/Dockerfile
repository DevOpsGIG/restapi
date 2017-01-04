FROM golang:1.8-alpine

COPY . src/github.com/devopsgig/restapi

WORKDIR src/github.com/devopsgig/restapi

RUN go build

CMD ["./restapi", "test"]
