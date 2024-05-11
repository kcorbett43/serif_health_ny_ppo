# syntax=docker/dockerfile:1
FROM golang:1.19

WORKDIR /myapp

COPY go.mod ./

RUN go mod download

COPY . .

EXPOSE 8080

CMD ["bash"]
