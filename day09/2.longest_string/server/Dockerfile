FROM golang:1.20 AS builder
WORKDIR /myapp
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./app ./main.go

FROM alpine:latest
WORKDIR /myapp
COPY --from=builder /myapp/app .
ENTRYPOINT [ "/myapp/app" ]