FROM iron/go:dev as builder

WORKDIR /go/src/github.com/vapor-ware/synse-amt-plugin
COPY . .

RUN make build


FROM iron/go
LABEL maintainer="vapor@vapor.io"

WORKDIR /plugin

COPY requirements.txt requirements.txt
RUN apk --update --no-cache add python2 \
    && apk --update --no-cache --virtual .build-dep add py2-pip \
    && pip install -r requirements.txt \
    && apk del .build-dep

COPY --from=builder /go/src/github.com/vapor-ware/synse-amt-plugin/build/plugin ./plugin
COPY config.yml .

CMD ["./plugin"]