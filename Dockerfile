# ------------------------------------------------------------------------------
# Build base image with all dependencies
# ------------------------------------------------------------------------------
FROM golang:1.12.5-alpine3.9 AS test_img

LABEL maintainer="Oliver Delgado <oliver.delgado@ibm.com>"

RUN apk update && apk upgrade && apk add --no-cache git
ENV APP_DIR=$GOPATH/src/prometheus-metrics/
RUN mkdir -p $APP_DIR
COPY . $APP_DIR
WORKDIR $APP_DIR
RUN go get
RUN go mod download
# Run command below to add GCC and dependencies
RUN apk add build-base
RUN CGO_ENABLED=1 GOOS=linux go build -installsuffix cgo -a -o /main .

# ------------------------------------------------------------------------------
# Create smaller runtime image
# ------------------------------------------------------------------------------
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=test_img /main ./
RUN chmod +x ./main
EXPOSE 8086
ENTRYPOINT ["./main"]