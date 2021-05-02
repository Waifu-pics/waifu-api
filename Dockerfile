FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]