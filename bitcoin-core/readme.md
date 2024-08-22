# bitcoin core docker

bitcoin core docker build

## 빌드 방법

`make` 명령어로 docker build 수행.

```
make
```

## ECR push

`make docker-push` 명령어로 docker push 실행

9154 계정의 ECR push 권한이 있는 AWS Credential 이 필요하다.

```
make docker-push
```

## Bitcoin core 로컬 실행

`make docker-run` 명령어로 bitcoin core 실행

실행하면 bitcoin block data 가 `$(PWD)/bitcoin_data` 경로에 저장되기 시작한다.

현재 날짜 2024-08-14 기준 약 600GB 이상의 용량이 필요하므로 용량주의 😵‍💫

동기화에 약 7일이 소요됨 😵

```
make docker-run
```

## Bitcoin core 실행 후 RPC 호출해보기

bitcoin.conf 에 아래와 같이 rpcuser, rpcpassword 가 설정되어있다.

```
rpcuser=abcrpcuser
rpcpassword=password
```

bitcoin core 를 실행한 다음 아래 명령어로 bitcoin RPC 를 사용해볼 수 있다.

현재 블록 높이 조회

```
curl -v -XPOST \
--header "Authorization: Basic YWJjcnBjdXNlcjpwYXNzd29yZA==" \
-d "{ \"jsonrpc\": \"1.0\", \"id\": \"1\", \"method\": \"getblockcount\", \"params\": [] }" \
--url "http://localhost:8332"
```
