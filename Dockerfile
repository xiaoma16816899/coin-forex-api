FROM golang:1.21-alpine AS build
WORKDIR /app

COPY .env /.env
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . ./
RUN go build -o /build

FROM alpine:latest

WORKDIR /

COPY --from=build /build /build
COPY --from=build .env /.env
EXPOSE 8000

ENTRYPOINT ["/build"]