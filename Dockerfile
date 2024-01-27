FROM golang:1.17-alpine3.14 as build

WORKDIR /code
COPY . /code/

WORKDIR /code/cmd/signup-server

RUN go build -o /web-server main.go


FROM alpine

WORKDIR /code

COPY --from=build  /web-server /bin/web-server

COPY --from=build  /code/assets /assets

ENTRYPOINT ["/bin/web-server"]