FROM golang:1.16 as builder
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app
FROM alpine:latest
COPY --from=builder /go/src/app/app /goapp/app
WORKDIR /goapp
COPY . /throwaway
RUN cp -r /throwaway/views ./views || echo "No resources to copy"
RUN rm -rf /throwaway
RUN apk --no-cache add ca-certificates
ENV PORT=8080
EXPOSE 8080
CMD ["/goapp/app"]