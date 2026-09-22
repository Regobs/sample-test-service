FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY app ./app
RUN CGO_ENABLED=0 go build -o /out/sample-test-service ./app

FROM alpine:3.20
RUN adduser -D -u 10001 app
USER app
COPY --from=build /out/sample-test-service /usr/local/bin/sample-test-service
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/sample-test-service"]
