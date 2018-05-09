FROM iron/go:dev as builder
WORKDIR /go/src/github.com/vapor-ware/synse-amt-plugin
COPY . .
RUN make build GIT_TAG="" GIT_COMMIT=""


FROM python:3.6-alpine
LABEL maintainer="vapor@vapor.io"

RUN apk add go --update-cache --repository http://dl-3.alpinelinux.org/alpine/edge/community/ --allow-untrusted

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

WORKDIR /plugin

COPY --from=builder /go/src/github.com/vapor-ware/synse-amt-plugin/build/plugin ./plugin
COPY config.yml .
COPY config/proto /etc/synse/plugin/config/proto

CMD ["./plugin"]