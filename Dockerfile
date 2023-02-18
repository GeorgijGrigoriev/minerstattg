FROM golang:1.19

WORKDIR /usr/src/minerstattg

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

ENV TOKEN=fillme

RUN go build -v -o /usr/local/bin/minerstat ./cmd/app/main.go

CMD ["minerstat"]