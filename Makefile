REVISION = "$(REVISION_NUM)\#$(shell date -u +%Y-%m-%d-%H:%M:%S)"

dbuild: 
	docker build -t dups . --build-arg  REVISION_NUM=$(REVISION)
build: 
	go mod download && go build -o dups -ldflags "-X 'github.com/msolimans/dups/pkg/build.Revision=$(REVISION)' -X 'github.com/msolimans/dups/pkg/build.Time=$(shell date)'" cmd/main.go
test: 
	go test -v -count=1 ./...
run:
	go run cmd/main.go