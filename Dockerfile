FROM golang:1.20.5-alpine3.18 as builder

RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /usr/src/app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download
COPY . .
RUN swag init
RUN go build -o ./bin/server .

FROM alpine

WORKDIR /usr/src/app
COPY --from=builder ["/usr/src/app/bin/server", "/usr/src/app/"]

EXPOSE 3000
CMD ["./server"]
