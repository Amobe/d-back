
all:

.PHONY: test
test:
	go test -count=1 -cover ./...

.PHONY: src.build
src.build:
	go build -ldflags="-s -w" -o d-back ./cmd/d-back/main.go
	go build -ldflags="-s -w" -o d-back-client ./cmd/d-back-client/main.go

.PHONY: docker.build
docker.build:
	docker build -t d-back -f Dockerfile.d-back . 
	docker build -t d-back-client -f Dockerfile.d-back-client . 
