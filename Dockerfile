##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN go build -o /cryar-api

##
## Deploy
##

FROM debian:buster-slim

WORKDIR /

COPY --from=build /cryar-api /cryar-api

EXPOSE 8080

ENTRYPOINT ["/cryar-api"]
