# Docker multi-stage build file
# Requires docker 17.05 or newer.

##
##
FROM golang:1.8 as builder

WORKDIR /go/src/docwhat.org/go-importd

COPY ./script/ ./script/
RUN ./script/bootstrap -u

COPY ./vendor/ ./vendor/
COPY ./*.go ./

RUN ./script/test -race
RUN ./script/lint
RUN CGO_ENABLED=0 govendor build -a -ldflags '-s -w' -o go-importd +local

##
##
FROM alpine:latest

ENV COLUMNS 80
EXPOSE 80

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/docwhat.org/go-importd/go-importd ./

ENTRYPOINT ["./go-importd"]

# vim: ft=dockerfile :
