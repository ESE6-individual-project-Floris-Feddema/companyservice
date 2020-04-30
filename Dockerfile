FROM golang:alpine AS builder

RUN apk update & apk add --no-cache git tzdata

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

ENV GIN_MODE=release

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /usr/local/bin/curl /curl

USER appuser:appuser

CMD ["/main"]