build:
	go build -o main

docker:
	docker build --no-cache --rm \
		--build-arg BUILD_TYPE=dev \
		--build-arg APP_VERSION=$(shell cat version.info) \
		--build-arg APP_NAME=bitcoin-core-api \
		-t bitcoin-core-api:$(shell cat version.info).dev .
	docker tag bitcoin-core-api:$(shell cat version.info).dev bitcoin-core-api:latest

docker-run:
	docker run -d --rm --env-file=.env --name bitcoin-core-api -p 3000:3000 bitcoin-core-api:$(shell cat version.info).dev

docker-stop:
	docker stop bitcoin-core-api