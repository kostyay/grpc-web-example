# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base
ADD . /src
WORKDIR /src
RUN go get ./...&& go build -o goapp

# final stage
FROM alpine
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app
COPY --from=build-env /src/goapp /app/
RUN chown -R appuser:appgroup /app
USER appuser
ENTRYPOINT ./goapp