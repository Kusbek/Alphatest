FROM golang:1.14.3-alpine3.11 AS builder
ENV PROJECT_REPO=alphatest
ENV APP_PATH=/go/src/${PROJECT_REPO}/
ENV GO111MODULE=on
WORKDIR ${APP_PATH}
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . ${APP_PATH}
RUN ls
RUN go build  -v ./cmd/apiserver

# stage 2
FROM alpine:3.11
ARG PROJECT_REPO
ENV PROJECT_REPO=alphatest
ENV APP_PATH=/go/src/${PROJECT_REPO}

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder ${APP_PATH}/apiserver ./

EXPOSE 8080
CMD ["./apiserver"]