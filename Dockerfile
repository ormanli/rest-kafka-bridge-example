FROM golang:1.15-alpine as builder
WORKDIR ..
RUN mkdir app
WORKDIR app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app github.com/ormanli/rest-kafka-bridge-example

FROM scratch
COPY --from=builder /app/app .
CMD ["./app"]
