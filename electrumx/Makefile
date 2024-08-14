docker-run:
	docker run -d \
	-v $(PWD)/electrum_data:/data \
	--rm --name electrumx \
	--network host \
	--env-file=.env \
	-p 8000:8000 -p 50001:50001 -p 50002:50002 lukechilds/electrumx:v1.16.0

docker-push:
	aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin 915486611144.dkr.ecr.ap-northeast-2.amazonaws.com
	docker tag lukechilds/electrumx:v1.16.0 915486611144.dkr.ecr.ap-northeast-2.amazonaws.com/lukechilds/electrumx:v1.16.0
	docker push 915486611144.dkr.ecr.ap-northeast-2.amazonaws.com/lukechilds/electrumx:v1.16.0