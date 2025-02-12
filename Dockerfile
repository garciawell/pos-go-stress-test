FROM golang:latest as builder
WORKDIR /app
COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o build main.go


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/build .
RUN mkdir -p /app/reports

ENV DISABLE_AUTO_OPEN=true

RUN apk --no-cache add ca-certificates

ENTRYPOINT ["./build"]