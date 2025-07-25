FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 4001

CMD ["./main"]