FROM golang:1.15.6-alpine

WORKDIR /go/src/github.com/agnjunio/go-rest-demo
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV PORT=80
EXPOSE 80

CMD ["go-rest-demo"]
