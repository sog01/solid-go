upES:
	docker-compose -f deployments/docker-compose.yaml up -d
downES:
	docker-compose -f deployments/docker-compose.yaml down
upESTest:
	docker-compose -f tests/build/docker-compose-test.yaml up -d
downESTest:
	docker-compose -f tests/build/docker-compose-test.yaml down
build:
	go build -o cmd/app-search/main cmd/app-search/main.go
ut:
	go test ./...
it: 
	go test ./tests/... -tags=integration -count=1
test:
	go test ./... -tags=integration -count=1 -coverprofile=coverage.out
run: build
	./cmd/app-search/main