FROM golang:1.20.5-alpine3.18 as builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .

RUN go build -o ./bin/server .

FROM alpine

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY --from=builder ["/usr/src/app/bin/server", "/usr/src/app/"]

EXPOSE 3000

CMD ["./server"]
