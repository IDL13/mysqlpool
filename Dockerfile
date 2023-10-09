FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o mysqlpool ./cmd/main.go

CMD ["./mysqlpool"]