FROM golang:1.17 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/greet ./greet.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/greet /app/greet
COPY --from=builder /app/etc /app/etc

EXPOSE 8888

CMD ["./greet", "-f", "etc/greet-api.yaml"]
