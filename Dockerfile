FROM golang:1.18

RUN go version

ENV GOPATH=/

COPY ./ ./

EXPOSE 8080:8080

RUN go mod download


RUN go build -o app ./cmd/main.go

CMD ["./app"]