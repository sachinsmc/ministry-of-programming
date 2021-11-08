# Building the binary of the App
FROM golang:1.17 AS build

WORKDIR /go/src/ministry-of-programming-tasks

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/ministry-of-programming-tasks/app .

COPY ./config/config.yml ./config/config.yml

EXPOSE 3003

CMD ["./app"]
