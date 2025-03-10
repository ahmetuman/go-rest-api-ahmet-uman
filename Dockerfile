FROM golang:1.23-alpine # Not sure abt version

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o app .

CMD ["/app/app"]
