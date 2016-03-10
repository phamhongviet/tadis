REDIS_TEST_IMAGE = redis:3-alpine
REDIS_TEST_PORT = 56379
REDIS_TEST_NAME = tadis_redis_test

GOLANG_IMAGE = golang:1.6-alpine

setup-test:
	docker inspect $(REDIS_TEST_NAME) &> /dev/null || \
	docker run -d -p $(REDIS_TEST_PORT):6379 --name $(REDIS_TEST_NAME) $(REDIS_TEST_IMAGE)

test: setup-test
	docker run --rm -v $(GOPATH):/go -v $(PWD):/app -w /app --link $(REDIS_TEST_NAME):redis $(GOLANG_IMAGE) go test

build:
	docker run --rm -v $(GOPATH):/go -v $(PWD):/app -w /app -e CGO_ENABLED=0 $(GOLANG_IMAGE) go build -ldflags "-s" -a -installsuffix cgo -o tadis

clean:
	docker stop $(REDIS_TEST_NAME)
	docker rm $(REDIS_TEST_NAME)
