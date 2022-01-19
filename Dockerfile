FROM golang:1.17.6
WORKDIR /cmd/signup-server

COPY . .

CMD ["go", "run", "/cmd/signup-server"]