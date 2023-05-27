FROM golang:1.20.4-alpine3.17 AS builder
ADD . /src
RUN apk add git
RUN cd /src && go build -o /src/bin/api -v /src/cmd/

# final stage
FROM alpine:3.17

EXPOSE 8080
ENV AOAI_RESOURCE_NAME=your-resource-name
ENV AOAI_DEPLOYMENT_NAME=your-deployment-name
ENV AOAI_API_VERSION=2023-05-15
ENV AOAI_API_KEY=your-api-key

WORKDIR /app
COPY --from=builder /src/bin/api /tmp/api
ENTRYPOINT /tmp/api