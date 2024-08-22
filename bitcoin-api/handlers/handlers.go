package handlers

import (
	"log"
	"net/http"

	"github.com/ahnlabio/bitcoin-core/bitcoin-api/service"
	"github.com/ahnlabio/bitcoin-core/bitcoin-api/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	btcSvc types.IBtcService
}

func NewHandler(apiService types.IBtcService) *Handler {
	return &Handler{
		btcSvc: apiService,
	}
}

// GetBalanceHandler godoc
// @Summary Show the balance of the address
// @Description get balance of the address
// @Tags balance
// @Accept  json
// @Produce  json
// @Param address query string true "address"
// @Success 200 {object} RootResponse
// @Router /v1/getBalance [get]
func (c Handler) GetBalanceHandler(ctx *gin.Context) {
	/*
		query parameter 로 입력받은 주소의 잔고를 조회한다.
		결과는 아래와 같은 json 형태로 반환한다.
		{
			"confirmd": 1000,
			"unconfirmd": 2000
		}

	*/
	address := ctx.Query("address")
	if address == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BadRequestErrorResp("address is required")})
		return
	}

	result, err := c.btcSvc.GetBalance(address)
	if err != nil {
		errResp(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c Handler) GetTransactionHandler(ctx *gin.Context) {
	txId := ctx.Query("txid")
	if txId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BadRequestErrorResp("txid is required")})
		return
	}

	if len(txId) != 64 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BadRequestErrorResp("txid must be of length 64")})
		return
	}

	result, err := c.btcSvc.GetTransaction(txId)
	if err != nil {
		errResp(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c Handler) GetUTXOHandler(ctx *gin.Context) {
	address := ctx.Query("address")
	if address == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BadRequestErrorResp("address is required")})
		return
	}

	result, err := c.btcSvc.GetUTXO(address)
	if err != nil {
		errResp(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c Handler) GetHistoryHandler(ctx *gin.Context) {
	address := ctx.Query("address")
	if address == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BadRequestErrorResp("address is required")})
		return
	}

	result, err := c.btcSvc.GetHistory(address)
	if err != nil {
		errResp(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func errResp(c *gin.Context, err error) {
	if errorInfo, ok := err.(*service.ServiceErr); ok {
		res := CommonErrorObject{
			Message: errorInfo.Msg,
			Text:    errorInfo.Text,
		}

		status := http.StatusInternalServerError
		if errorInfo.Text == service.ERR_INVALID_ADDRESS {
			status = http.StatusBadRequest
		}

		log.Printf("[ERROR] err: %s, url: %s, status: %d\n", err.Error(), c.Request.URL, status)
		c.JSON(status, gin.H{"error": &res})
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
