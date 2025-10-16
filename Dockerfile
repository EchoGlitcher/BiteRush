FROM golang:1.25.3-alpine3.21 AS build

RUN apk add make git
WORKDIR /build
COPY . .

RUN make build

FROM alpine:3.21

WORKDIR /sv
COPY --from=build /build/.builds/biterush ./
ENTRYPOINT ["./biterush"]
