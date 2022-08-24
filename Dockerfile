ARG GO_VERSION=1.14.0
FROM golang:${GO_VERSION}-buster AS builder
RUN apt-get update && apt-get install -y make git

WORKDIR /frappe

ARG CI_BUILD_TOKEN
ENV CI_BUILD_TOKEN ${CI_BUILD_TOKEN}

# RUN make build

VOLUME ["/etc/frappe"]
EXPOSE 8888
CMD ["/bin/bash"]
