FROM golang

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./tmp/main ./cmd

EXPOSE 8080
CMD ["/app/tmp/main"]