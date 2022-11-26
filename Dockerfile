FROM golang:1.18-alpine

ENV GO111MODULE=on

RUN mkdir /app
ADD . /app/

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080
CMD ["/app/main"]