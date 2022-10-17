build: 
	docker build -t dups . --build-arg  REVISION_NUM=sha
test: 
	go test -v ./...
run:
	go run cmd/main.go