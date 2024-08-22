# electrum

docs

https://electrumx-spesmilo.readthedocs.io/en/latest/

## electrumx 로컬 실행

`make docker-run` 명령어로 electrumx core 실행

실행하면 electrum index data 가 `$(PWD)/electrum_data` 경로에 저장되기 시작한다.

현재 날짜 2024-08-14 기준 약 100GB 이상의 용량이 필요하므로 용량주의 😵‍💫

동기화에 약 7일이 소요됨 😵

동기화가 완료(bitcoin core node 와 높이가 같아짐)되기 전까지 외부 RPC 요청이 허용되지 않음.

electrum_rpc 내부 명령어만 사용 가능(문서 참고)

```
make docker-run
```

## 참고 링크

### github

https://github.com/lukechilds/docker-electrumx

### docker hub

https://hub.docker.com/r/lukechilds/electrumx
