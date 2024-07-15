FROM golang:alpine as builder
WORKDIR /go/src/K8sVisualization/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="lsx@lsx.com"
WORKDIR /go/src/K8sVisualization/server

COPY --from=0 /go/src/K8sVisualization/conf ./conf 
COPY --from=0 /go/src/K8sVisualization/server ./server 
EXPOSE 8889
ENTRYPOINT ./server