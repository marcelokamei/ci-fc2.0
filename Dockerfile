FROM golang:alpine3.13 as builder

WORKDIR /app

COPY . .


RUN go build -o math

FROM scratch

COPY --from=builder /app .

CMD ["./math"]