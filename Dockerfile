FROM golang:1.20.2-alpine as builder
WORKDIR /opt
COPY . .
RUN go mod tidy; \
    go fmt ./...; \
    go mod download;
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./unnecessarywebserver .

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /opt/unnecessarywebserver /app/unnecessarywebserver
EXPOSE 8080
CMD [ "./unnecessarywebserver" ]