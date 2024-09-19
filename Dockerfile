FROM golang:1.22.1-bookworm AS os


FROM os AS dev
RUN apt-get update -qq -y \
    && apt-get install -y build-essential ca-certificates file make gcc 
RUN   mkdir /go/pkg
RUN   chmod a+rwx /go/pkg
ARG APP_USER_ID
RUN useradd -s /bin/sh -u ${APP_USER_ID} -m appuser
USER appuser
WORKDIR /go/src/app

FROM dev AS builder
COPY . /go/src/app
RUN go build -buildvcs=false

FROM os AS prod
COPY --from=builder /go/src/app/bin/setkafka /usr/bin/setkafka
COPY --from=builder /go/src/app/setkafka.yaml /etc/setkafka/setkafka.yaml
