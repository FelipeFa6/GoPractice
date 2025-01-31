APP_NAME=bankapp
CONTAINER_NAME=pg_container_bank
HOST_PORT=5432
CONTAINER_PORT=5432

install:
	docker build -t $(APP_NAME) .
	docker run -d --rm --name $(CONTAINER_NAME) -p $(HOST_PORT):$(CONTAINER_PORT) $(APP_NAME)

run:
	@go run .

build:
	go build -o $(APP_NAME)

clean:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true
	docker rmi $(APP_NAME) || true
	@if [ -f $(APP_NAME) ]; then rm $(APP_NAME); fi

default: run

