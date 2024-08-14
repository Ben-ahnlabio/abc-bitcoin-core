# electrum-api

electrum api

## elemctrum api 서버

`github.com/checksum0/go-electrum` 를 이용하여 electrumx RPC 에 요청을 보내는 HTTP API Server

## API Reference

### Get Balance

비트코인 주소의 잔액을 가져온다

```http
GET /v1/balance
```

| Query Parameter | Type     | Description                   |
| :-------------- | :------- | :---------------------------- |
| `address`       | `string` | **Required**. Bitcoin Address |

**Response Example**

```json
{
  "address": "bc1qanfh6n9csne5swjer6wmd2djugcy5y6eqtws67",
  "confirmed": 1000,
  "unconfirmed": 0
}
```

### Get Transaction

비트코인 transaction 정보를 가져온다

```http
GET /v1/transaction
```

| Query Parameter | Type     | Description                          |
| :-------------- | :------- | :----------------------------------- |
| `txid`          | `string` | **Required**. Bitcoin transaction id |

**Response Example**

```json
{
  "block_hash": "0000000000000000000197cc7a6d13c63cd086ca46372cc94577ff62fa56d568",
  "tx_hash": "72c9cde2e7de731ab3b517bc582d23a91b0d2714ed1fdd96e580e6446a236a0f",
  "confirmations": 10
}
```

### Get UTXO

비트코인 주소 UTXO 를 가져온다

```http
GET /v1/utxo
```

| Query Parameter | Type     | Description                   |
| :-------------- | :------- | :---------------------------- |
| `address`       | `string` | **Required**. Bitcoin Address |

**Response Example**

```json
{
  "address": "bc1qanfh6n9csne5swjer6wmd2djugcy5y6eqtws67",
  "UTXOs": [
    {
      "height": 856675,
      "tx_pos": 1,
      "tx_hash": "72c9cde2e7de731ab3b517bc582d23a91b0d2714ed1fdd96e580e6446a236a0f",
      "value": 21094
    }
  ]
}
```
