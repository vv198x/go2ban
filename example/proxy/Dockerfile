FROM golang:1.19.3-alpine as build

WORKDIR /usr/local/go/src/proxy/
ADD . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -v -a -o proxy -ldflags="-w -s"

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/local/go/src/proxy/proxy /

ENTRYPOINT ["./proxy"]