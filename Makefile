.PHONY: up down build run rm

# IMAGE_NAME represents the name of image build from Dockerfile
IMAGE_NAME=user-crud

# CONTAINER_NAME represents the name of the docker container that runs user-crud server
CONTAINER_NAME=pawelWritesCode.${IMAGE_NAME}.server

# up turns up containers as defined in compose.yaml file
up:
	docker compose up --force-recreate --build -d && docker ps -a

# down turns down containers defined in compose.yaml file
down:
	docker compose down

# logs show last 100 lines of logs to stdout from server container
logs:
	docker logs -n 100 -f ${CONTAINER_NAME}

# build builds image as defined in Dockerfile
build:
	docker build -t ${IMAGE_NAME} .

# run runs server's docker image
run:
	docker run -dp 1234:1234 --name ${CONTAINER_NAME} ${IMAGE_NAME} || docker ps -a

# rm stops & removes server's docker image
rm:
	docker rm -f ${CONTAINER_NAME}