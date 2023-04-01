package controller

import (
	"net/http"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type processController struct {
	ProcessService service.ProcessService
}

func NewProcessController(processService service.ProcessService) processController {
	return processController{ProcessService: processService}
}

func (c *processController) AcceptOrder(ctx *gin.Context) {
	driverId, _ := strconv.Atoi(ctx.Param("driverId"))
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))

	res, err := c.ProcessService.AcceptOrder(driverId, orderId)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in accept order service."})
		return
	}

	ctx.JSON(200, res)
}

func (c *processController) CancellOrder(ctx *gin.Context) {
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))

	if res, err := c.ProcessService.CancellOrder(orderId); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in cancell order service.", "res": res})
		return
	}

	ctx.JSON(200, "Cancelled.")
}

func (c *processController) GetOrdersInProcessDriver(ctx *gin.Context) {
	driverId, _ := strconv.Atoi(ctx.Param("driverId"))

	res, err := c.ProcessService.GetOrdersInProcessDriver(driverId)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in GetOrdersInProcessDriver order service."})
		return
	}

	ctx.JSON(200, res)
}

func (c *processController) GetOrdersInProcessUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("userId"))

	res, err := c.ProcessService.GetOrdersInProcessDriver(userId)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in GetOrdersInProcessUser order service."})
		return
	}

	ctx.JSON(200, res)
}

func (c *processController) FinishOrder(ctx *gin.Context) {
	driverId, _ := strconv.Atoi(ctx.Param("driverId"))
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))

	res, err := c.ProcessService.FinishOrder(driverId, orderId)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in finish order service."})
		return
	}

	ctx.JSON(200, res)
}
