ARG GO_VERSION=1.14.0
FROM golang:${GO_VERSION}-buster AS builder
RUN apt-get update && apt-get install -y make git

WORKDIR /frappe

ARG CI_BUILD_TOKEN
ENV CI_BUILD_TOKEN ${CI_BUILD_TOKEN}
RUN git config --global url."https://gitlab-ci-token:$CI_BUILD_TOKEN@gitlab.zerodha.tech".insteadOf "https://gitlab.zerodha.tech"

ENV CGO_ENABLED=0 GOOS=linux GOSUMDB="sum.golang.org" GOPROXY="https://goproxy.zerodha.tech" GONOSUMDB="*.zerodha.tech/*"

# RUN make build

VOLUME ["/etc/frappe"]
EXPOSE 8888
# CMD ["./rand.bin", "--config", "/etc/frappe/config.toml"]
CMD ["/bin/bash"]
