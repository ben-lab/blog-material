FROM golang as build
COPY reverse-proxy.go /go/src/
RUN go get github.com/tv42/httpunix
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' src/reverse-proxy.go

FROM scratch
COPY --from=build /go/reverse-proxy /bin/reverse-proxy
USER root
ENTRYPOINT ["/bin/reverse-proxy"]
