FROM golang:1.17
WORKDIR /code
COPY . /code/
WORKDIR /code/cmd/signup-server
CMD ["go","run","main.go"]