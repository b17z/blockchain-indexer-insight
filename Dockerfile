FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev

RUN go get github.com/kelseyhightower/envconfig
RUN go get github.com/googollee/go-socket.io
RUN go get github.com/googollee/go-engine.io
RUN go get github.com/gorilla/websocket
RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/gertjaap/blockchain-indexer-insight
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["blockchain-indexer-insight"]