FROM golang:1.15.8-alpine3.13 as build-env
LABEL maintainer="Arpeet Gupta <arpeet.gupta96@gmail.com>"

RUN mkdir /product
WORKDIR /product
COPY go.mod . 
COPY go.sum .

RUN --mount=type=cache,target=/go go mod download
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -installsuffix cgo -o /opt/product cmd/server/*.go

FROM alpine:3.13.1
# RUN apk add --no-cache bash && \
# RUN   apk add --update --no-cache ca-certificates git
COPY --from=build-env /opt/product /go/bin/product
ENTRYPOINT ["/go/bin/product"]
CMD ["--port", "8080"]
# Add `CMD` and Expose port for REST Endpoint 
EXPOSE 8080
