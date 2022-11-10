upES:
	docker-compose -f deployments/docker-compose.yaml up -d
downES:
	docker-compose -f deployments/docker-compose.yaml down
build:
	go build -o cmd/app-search/main cmd/app-search/main.go
run: build
	./cmd/app-search/main