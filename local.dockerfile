FROM alpine
MAINTAINER Christian Höltje <docwhat@gerf.org>

ENV COLUMNS 80
EXPOSE 80

RUN apk add --no-cache ca-certificates
COPY ["go-importd_linux_amd64", "/go-importd"]

ENTRYPOINT ["/go-importd"]

# vim: ft=dockerfile :
