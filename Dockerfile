FROM golang:1.18-alpine

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

CMD [ "go", "run", "main.go" ]
