FROM golang:1.8-alpine
ADD . /go/src/hello-app
RUN apk update && apk add git
RUN go get -u github.com/lib/pq
RUN go install hello-app

FROM alpine:latest
COPY --from=0 /go/bin/hello-app .
ENV PORT 8080
CMD ["./hello-app"]
