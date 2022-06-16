FROM alpine:latest

ARG BUILD_APP_NAME

ENV PROJECT_DIR=/go

RUN apk add tzdata

COPY ./build/${BUILD_APP_NAME} ${GOBIN}/app
COPY ./configs ${PROJECT_DIR}/configs

CMD ${PROJECT_DIR}/bin/app -c ${PROJECT_DIR}/configs/config.yml
