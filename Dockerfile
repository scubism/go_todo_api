FROM golang:1.5.1

MAINTAINER scubism

# set main app directory
ENV APP_SRC_DIR="/usr/src/app"
WORKDIR ${APP_SRC_DIR}
VOLUME [${APP_SRC_DIR}]

# set go path settings
# TODO ENV GO15VENDOREXPERIMENT 1
ENV GOPATH $GOPATH:${APP_SRC_DIR}/_vendor
ENV PATH $PATH:${APP_SRC_DIR}/_vendor/bin

# set common go settings
RUN go get github.com/mattn/gom \
  && cd /go/src/github.com/mattn/gom \
  && git reset --hard 78a909167da6e3b7ea010766e09738d086427f6d \
  && cd ${APP_SRC_DIR} \
  && go get github.com/mattn/gom
ENV GIN_MODE "release"

# === app specific settings ===
COPY . ${APP_SRC_DIR}
RUN gom install

ENV MONGO_HOST "127.0.0.1"
ENV MONGO_PORT "27017"
ENV MONGO_DB "go_todo_api"

EXPOSE 3000

ENTRYPOINT ["./docker-entrypoint.sh"]
