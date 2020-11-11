FROM golang:latest

WORKDIR /api
COPY . .
ENV GO111MODULE=on

RUN go get github.com/labstack/echo/v4 && \
    go get github.com/pilu/fresh
# go mod init go-tools && \


CMD ["go", "run", "main.go"]
# CMD ["fresh"]
