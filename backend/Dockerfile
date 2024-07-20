FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

COPY --from=builder /app/tables ./tables

EXPOSE 8080

ENV GIN_MODE=release

CMD ["./main"]