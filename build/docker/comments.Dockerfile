FROM golang:alpine as build

COPY ./ /app
WORKDIR /app

RUN apk update && apk upgrade
RUN apk add --update build-base libwebp-dev
RUN go build -o main.out ./internal/comment/server/comment_app.go ./internal/comment/server/comment_server.go

FROM alpine


RUN apk update && apk upgrade
RUN apk add --update libwebp-dev
COPY --from=build /app/main.out /
COPY --from=build /app/configs/. /configs/

ENTRYPOINT ./main.out