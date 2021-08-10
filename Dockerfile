# Builder
FROM golang:1.16-alpine as builder

WORKDIR /app

RUN apk update && apk upgrade && apk --update add git make

COPY . .

RUN make prometheus-converter

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8080

COPY --from=builder /app/prometheus-converter /app
COPY --from=builder /app/config.yaml /app
COPY --from=builder /app/static/ /app/static/

USER 65532:65532

ENTRYPOINT ["/app/prometheus-converter"]

