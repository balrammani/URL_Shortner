# Use golang-based image for container
FROM golang:1.16

WORKDIR /go/src/rest-url-shortener

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["go", "run", "./cmd/app/main.go"]