FROM golang:1.11

WORKDIR /api
COPY . .
ENV GO111MODULE=on

RUN go get github.com/labstack/echo/v4 && \
    go get github.com/pilu/fresh

CMD ["fresh"]
