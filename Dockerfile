# Builder Image
FROM vaporio/golang:1.11 as builder
WORKDIR /go/src/github.com/vapor-ware/synse-amt-plugin
COPY . .

# If the vendored dependencies are not present in the docker build context,
# we'll need to do the vendoring prior to building the binary.
RUN if [ ! -d "vendor" ]; then make dep; fi
RUN make build CGO_ENABLED=0


FROM vaporio/python:3.6-slim
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/vapor-ware/synse-amt-plugin/build/plugin ./plugin

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

# Add in default plugin configuration.
COPY config.yml /etc/synse/plugin/config/config.yml

# Image Metadata -- http://label-schema.org/rc1/
# This should be set after the dependency install so we can cache that layer
ARG BUILD_DATE
ARG BUILD_VERSION
ARG VCS_REF

LABEL maintainer="vapor@vapor.io" \
      org.label-schema.schema-version="1.0" \
      org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="vaporio/amt-plugin" \
      org.label-schema.vcs-url="https://github.com/vapor-ware/synse-amt-plugin" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vendor="Vapor IO" \
      org.label-schema.version=$BUILD_VERSION

ENTRYPOINT ["./plugin"]