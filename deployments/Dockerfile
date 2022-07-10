FROM alpine:latest

ARG BUILD_APP_NAME

ENV PROJECT_DIR=/go

RUN apk add tzdata

COPY ./build/${BUILD_APP_NAME} ${PROJECT_DIR}/bin/app
COPY ./configs ${PROJECT_DIR}/configs

CMD ${PROJECT_DIR}/bin/app -c ${PROJECT_DIR}/configs/config.yml
