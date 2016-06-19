FROM golang:1.6.2
MAINTAINER Dmitry Moskowski <me@corpix.ru>

WORKDIR /go/src
RUN apt-get update && apt-get install -y --force-yes libsasl2-dev
RUN go get github.com/corpix/geochats-backend/... \
    && cd github.com/corpix/geochats-backend \
    && go get -u github.com/tools/godep \
    && godep restore \
    && go build -o /usr/bin/geochats-backend geochats-backend/main.go \
    && cd /go && rm -rf src

EXPOSE 3000

CMD ["/usr/bin/geochats-backend", "serve"]
