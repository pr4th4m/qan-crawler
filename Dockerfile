# # Build image
# docker build -t "qan-crawler:dev" .
#
# # Run container
# docker run --name qan-crawler -p 8080:8080 -d qan-crawler:dev
# docker logs -f qan-crawler
FROM golang:1.10.0-stretch

# Copy binary and config file
WORKDIR /go/src/github.com/pratz/qan-crawler
COPY . .

# Install app dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# Build binary
RUN go build

# Run app
EXPOSE 8080
ENTRYPOINT ["./qan-crawler"]
