FROM alpine:latest AS certificates
RUN apk add --no-cache ca-certificates

FROM scratch AS release
ENV COLUMNS 80
EXPOSE 80
COPY --from=certificates /etc/ssl/ /etc/ssl/
COPY go-importd /go-importd
ENTRYPOINT ["/go-importd"]
LABEL \
  "org.label-schema.name"="go-importd" \
  "org.label-schema.description"="A 'go get' redirector for vanity names." \
  "org.label-schema.url"="https://github.com/docwhat/go-importd" \
  "org.label-schema.vcs-url"="https://github.com/docwhat/go-importd.git" \
  "org.label-schema.schema-version"="1.0"
