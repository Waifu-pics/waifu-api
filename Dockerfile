FROM golang:alpine
COPY . .
RUN go build -o app

ENTRYPOINT ./app