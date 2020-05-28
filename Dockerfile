FROM golang:alpine AS builder

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go mod download
RUN go mod verify

##ENV GIN_MODE=release

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#RUN go build -o main .

EXPOSE 80

CMD ["./main"]