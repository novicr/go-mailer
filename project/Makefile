
MAILER_BINARY=mailer

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
build: build_mailer
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_mailer: builds mailer application
build_mailer: test
	@echo "Building mailer binary..."
	cd ../mailer && CGO_ENABLED=0 go build -o ${MAILER_BINARY} .
	@echo "Done!"

test: 
	@echo "Testing mailer application"
	cd ../mailer && go test ./...
	@echo "Done!"
