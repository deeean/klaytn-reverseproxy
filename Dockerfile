FROM golang:1.17-alpine

WORKDIR /app

ENV USERNAME $USERNAME \
    PASSWORD $PASSWORD \
    CYPRESS_RPC_URL $CYPRESS_RPC_URL \
    BAOBAB_RPC_URL $BAOBAB_RPC_URL

WORKDIR /go/src/app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

ENTRYPOINT ["./main"]