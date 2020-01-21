FROM golang:alpine AS builder

WORKDIR $GOPATH/src/0xdiba/go-service/
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/go-service

FROM scratch
COPY --from=builder /go/bin/go-service /go/bin/go-service

EXPOSE 8000

ENTRYPOINT ["/go/bin/go-service"]
