
FROM docker.io/vaporio/python:3.6-slim

LABEL org.label-schema.schema-version="1.0" \
      org.label-schema.name="vaporio/amt-plugin" \
      org.label-schema.vcs-url="https://github.com/vapor-ware/synse-amt-plugin" \
      org.label-schema.vendor="Vapor IO"

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

# Add in default plugin configuration.
COPY config.yml /etc/synse/plugin/config/config.yml

# Add plugin scripts.
COPY scripts scripts

# Copy the executable.
COPY synse-amt-plugin ./plugin

ENTRYPOINT ["./plugin"]
