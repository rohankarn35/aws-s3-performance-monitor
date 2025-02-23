# Makefile

# Variables
DOCKER_IMAGE_NAME=aws-golang
DOCKER_CONTAINER_NAME=aws-golang-container
ENV_FILE=.env

# Build the Docker image
build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# Run the Docker container
run: build
	docker run -d --env-file $(ENV_FILE) --name $(DOCKER_CONTAINER_NAME) -p 8080:8080 $(DOCKER_IMAGE_NAME)

# Stop and remove the Docker container
stop:
	docker stop $(DOCKER_CONTAINER_NAME) || true
	docker rm $(DOCKER_CONTAINER_NAME) || true

# Clean up Docker images and containers
clean: stop
	docker rmi $(DOCKER_IMAGE_NAME) || true

# Rebuild and run the Docker container
rerun: clean run