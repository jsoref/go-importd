FROM alpine
MAINTAINER Christian Höltje <docwhat@gerf.org>

ENV GO_IMPORTD_VERSION 0.1.0
ENV COLUMNS           80
EXPOSE 80

RUN apk add --no-cache ca-certificates
ADD ["https://github.com/docwhat/go-importd/releases/download/${GO_IMPORTD_VERSION}/go-importd_linux_amd64", "/go-importd"]
RUN ["chmod", "0755", "/go-importd"]

ENTRYPOINT ["/go-importd"]

# vim: ft=dockerfile :
