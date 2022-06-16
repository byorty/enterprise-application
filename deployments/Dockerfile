FROM alpine:latest

ARG BUILD_APP_NAME

ENV GOPATH=/go
ENV GOBIN=${GOPATH}/bin

RUN apk add tzdata

COPY ./build/${BUILD_APP_NAME} ${GOBIN}/app
COPY ./configs ${GOPATH}/configs

CMD ${GOBIN}/app -c ${GOPATH}/configs/config.yml
