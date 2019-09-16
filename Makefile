.PHONY: build     # Builds docker image for deployment
build:
	docker build -t curve-tech-test-go:latest .

.PHONY: up        # creates and runs the container from built image
up:
	docker run --rm -itd --name curve-tech-test-hadi -p 8091:8091 --env-file .env curve-tech-test-go:latest

.PHONY: down      # Stops the Docker containers
down:
	docker stop curve-tech-test-hadi