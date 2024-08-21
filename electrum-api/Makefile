build:
	go build -o main

docker:
	docker build --no-cache --rm \
		--build-arg APP_VERSION=$(shell cat version.info) \
		--build-arg APP_NAME=bitcoin-core-api \
		-t bitcoin-core-api:$(shell cat version.info).dev .
	docker tag bitcoin-core-api:$(shell cat version.info).dev bitcoin-core-api:latest

docker-push: init
	aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin 915486611144.dkr.ecr.ap-northeast-2.amazonaws.com
	docker tag bitcoin-core-api:latest 915486611144.dkr.ecr.ap-northeast-2.amazonaws.com/bitcoin-core-api:$(shell cat version.info).dev
	docker push 915486611144.dkr.ecr.ap-northeast-2.amazonaws.com/bitcoin-core-api:$(shell cat version.info).dev

docker-run:
	docker run -d --rm --env-file=.env --name bitcoin-core-api -p 3000:3000 bitcoin-core-api:$(shell cat version.info).dev
	docker logs -f bitcoin-core-api

docker-stop:
	docker stop bitcoin-core-api