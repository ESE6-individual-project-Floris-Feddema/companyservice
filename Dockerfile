FROM golang:alpine AS builder

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

#ENV GIN_MODE=release

COPY --from=builder /app/main .
COPY --from=builder /app/config.env .
CMD ["/main"]