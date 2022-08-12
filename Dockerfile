FROM golang:1.18

WORKDIR /app


COPY src/go.mod ./
COPY src/go.sum ./


RUN go mod download

COPY src/cmd ./cmd
COPY src/load ./load


RUN go build  /app/load/main.go

CMD [ "/app/main"]