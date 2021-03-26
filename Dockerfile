FROM golang:alpine as builder

RUN go get github.com/eslizn/docker-plugin-volume-cos/...

FROM epurs/cosfs

COPY --from=builder /go/bin/docker-plugin-volume-cos /bin/
