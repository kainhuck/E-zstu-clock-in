ARG BUILDER_BASE=golang:1.15-alpine3.12
FROM ${BUILDER_BASE} AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --update --no-cache make git build-base gcc

WORKDIR /ezstu

COPY . .
RUN go mod download
RUN make build

FROM alpine:3.12

WORKDIR /ezstu

COPY --from=builder /ezstu/card .
COPY --from=builder /ezstu/config.json .

ENTRYPOINT ["./card"]