package handlers

import (
	"log"
	"net/http"

	"github.com/ahnlabio/bitcoin-core/electrum-api/container"
	"github.com/ahnlabio/bitcoin-core/electrum-api/service"
	"github.com/gin-gonic/gin"
)

// GetBalanceHandler godoc
// @Summary Show the balance of the address
// @Description get balance of the address
// @Tags balance
// @Accept  json
// @Produce  json
// @Param address query string true "address"
// @Success 200 {object} RootResponse
// @Router /v1/getBalance [get]
func GetBalanceHandler(c *gin.Context) {
	/*
		query parameter 로 입력받은 주소의 잔고를 조회한다.
		결과는 아래와 같은 json 형태로 반환한다.
		{
			"confirmd": 1000,
			"unconfirmd": 2000
		}

	*/
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": BadRequestErrorResp("address is required")})
		return
	}

	conatiner := container.GetInstnace()
	service := conatiner.GetBitcoinApiService()
	result, err := service.GetBalance(address)
	if err != nil {
		errResp(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetTransactionHandler(c *gin.Context) {
	txId := c.Query("txid")
	if txId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": BadRequestErrorResp("txid is required")})
		return
	}

	conatiner := container.GetInstnace()
	service := conatiner.GetBitcoinApiService()
	result, err := service.GetTransaction(txId)
	if err != nil {
		errResp(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetUTXOHandler(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": BadRequestErrorResp("address is required")})
		return
	}

	conatiner := container.GetInstnace()
	service := conatiner.GetBitcoinApiService()
	result, err := service.GetUTXO(address)
	if err != nil {
		errResp(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetHistoryHandler(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": BadRequestErrorResp("address is required")})
		return
	}

	conatiner := container.GetInstnace()
	service := conatiner.GetBitcoinApiService()
	result, err := service.GetHistory(address)
	if err != nil {
		errResp(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

type GetUTXOResponse struct {
	Address string  `json:"address"`
	UTXOs   []*UTXO `json:"utxos"`
}

type UTXO struct {
	Height   uint32 `json:"height"`
	Position uint32 `json:"tx_pos"`
	Hash     string `json:"tx_hash"`
	Value    uint64 `json:"value"`
}

type GetBalanceResponse struct {
	Address    string `json:"address"`
	Confirmd   int    `json:"confirmd"`
	Unconfirmd int    `json:"unconfirmd"`
}

type GetTransactionResponse struct {
	BlockHash     string `json:"block_hash"`
	TxHash        string `json:"tx_hash"`
	Confirmations int    `json:"confirmations"`
}

type GetHistoryResponse struct {
	Address   string     `json:"address"`
	Histories []*History `json:"histories"`
}

type History struct {
	Height int32  `json:"height"`
	TxHash string `json:"tx_hash"`
	Fee    uint32 `json:"fee,omitempty"`
}

func errResp(c *gin.Context, err error) {
	log.Printf("[ERROR] err: %s, url: %s\n", err.Error(), c.Request.URL)
	if errorInfo, ok := err.(*service.ServiceErr); ok {
		log.Printf("errorInfo: %v\n", errorInfo)
		res := CommonErrorObject{
			Message: errorInfo.Msg,
			Text:    errorInfo.Text,
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": &res})
		return
	}
	res := CommonErrorObject{
		Message: err.Error(),
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": &res})
}

type CommonErrorObject struct {
	Text    string `json:"text" example:""`
	Message string `json:"message" example:""`
}
