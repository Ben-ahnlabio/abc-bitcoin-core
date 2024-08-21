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
  "utxos": [
    {
      "height": 853533,
      "tx_pos": 0,
      "tx_hash": "9870aa7720c577d5e3f01de8d05bda1356b725847a7eae5c0d81618a8f8d9e28",
      "value": 2000
    }
  ]
}
```

### Get History

비트코인 주소 history 를 가져온다

```http
GET /v1/history
```

| Query Parameter | Type     | Description                   |
| :-------------- | :------- | :---------------------------- |
| `address`       | `string` | **Required**. Bitcoin Address |

**Response Example**

```json
{
  "address": "bc1qanfh6n9csne5swjer6wmd2djugcy5y6eqtws67",
  "histories": [
    {
      "height": 852690,
      "tx_hash": "60d14e38db4a6bd2a42c26c33479d1532d656ff314461168346839c86b60dde3"
    },
    {
      "height": 853533,
      "tx_hash": "894af5cb799532ca63f46decabe418a6df70a4942e678e572c309c578d5eaab7"
    },
    {
      "height": 853533,
      "tx_hash": "9870aa7720c577d5e3f01de8d05bda1356b725847a7eae5c0d81618a8f8d9e28"
    }
  ]
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
  "block_hash": "000000000000000000018824e004cc5e283f383bee364cd71322288c11ea2cb6",
  "tx_hash": "84b28ac934d0ac882e3d40e691778cf915202b38c93d018726631984114d3859",
  "confirmations": 4155
}
```
