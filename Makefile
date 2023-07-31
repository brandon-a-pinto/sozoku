build_server:
	@go build -C cmd/server -o bin/server

build_client:
	@go build -C cmd/client -o bin/client

run_server: build_server
	@./cmd/server/bin/server

run_client: build_client
	@./cmd/client/bin/client
