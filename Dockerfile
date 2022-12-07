FROM golang:1.16 as builder
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app
FROM alpine:latest
COPY --from=builder /go/src/app/app /goapp/app
WORKDIR /goapp
COPY ./views /goapp/views
COPY ./assets /goapp/assets
RUN apk --no-cache add ca-certificates
ENV PORT=8080
EXPOSE 8080
CMD ["/goapp/app"]