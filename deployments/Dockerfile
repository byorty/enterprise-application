FROM alpine:latest

ARG BUILD_APP_NAME

ENV GOPATH=/go
ENV GOBIN=${GOPATH}/bin

RUN apk add tzdata

COPY ./build/${BUILD_APP_NAME} ${GOBIN}/app

CMD ${GOBIN}/app
