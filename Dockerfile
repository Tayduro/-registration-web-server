FROM golang:1.17.6
WORKDIR /cmd/signup-server

COPY . .

CMD ["go", "run", "/code/cmd/signup-server/main.go"]