FROM golang:1.22
WORKDIR /go/src/app
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]