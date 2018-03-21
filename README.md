**This is a test project**

#### qan-crawler
Crawl web pages recursively


#### Install and run with Docker (Recommended)

    docker build -t "qan-crawler:dev" .
    docker run --name qan-crawler -p 8080:8080 -d qan-crawler:dev

    # Check logs
    docker logs -f qan-crawler


#### Install and run manually

    # Install dependency manager
    go get -u github.com/golang/dep/cmd/dep
    dep ensure

    # Run unit tests
    go test ./...

    # Build and install
    go build
    ./qan-crawler


#### Usage

    curl '0.0.0.0:8080/crawl'

    # Url param
    curl '0.0.0.0:8080/crawl?url=https://www.qantasmoney.com'

    # Depth param
    curl '0.0.0.0:8080/crawl?url=https://www.qantasmoney.com&depth=2'

    # Exclude uri param
    http '0.0.0.0:8080/crawl?url=https://www.qantasmoney.com&depth=2&exclude=https://www.qantasstore.com.au/

    # Defaults
	url = "https://google.com"
	depth = 1
    exclude = ["/", "#"]


#### Note
- Depending on depth settings some urls may take more than 30secs, please set timeout accordingly.
- Actually, this project should have a advanced caching and status polling api, however, leaving that off for now as its a test project.
