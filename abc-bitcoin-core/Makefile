init:
	docker build -t abc-bitcoin-core .

docker-push:
	aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin 915486611144.dkr.ecr.ap-northeast-2.amazonaws.com
	docker tag abc-bitcoin-core:latest 915486611144.dkr.ecr.ap-northeast-2.amazonaws.com/abc-bitcoin-core:$(shell cat version.info)
	docker push 915486611144.dkr.ecr.ap-northeast-2.amazonaws.com/abc-bitcoin-core:$(shell cat version.info)

docker-stop:
	docker stop abc-bitcoin-core

docker-run:
	docker run -d --rm --name abc-bitcoin-core -p 28332:28332 -p 8332:8332 --network bitcoin-core -v $(PWD)/bitcoin_data:/root/bitcoin_data abc-bitcoin-core:latest
	docker logs -f abc-bitcoin-core