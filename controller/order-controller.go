package controller

import (
	"fmt"
	"net/http"
	"qkeruen/dto"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	OrderService service.OrderService
}

func NewOrderController(service service.OrderService) orderController {
	return orderController{OrderService: service}
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var order dto.OrderRequest

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit
	}

	if err := c.OrderService.CreateOrder(id, order); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in create order service."})
		return
	}

	ctx.JSON(200, "Order created")
}

func (c *orderController) GetOrders(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res, err := c.OrderService.GetOrders(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in search order service."})
		return
	}

	ctx.JSON(200, res)
}

func (c *orderController) GetMyOrders(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := c.OrderService.GetMyOrders(id)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in get my order service."})
		return
	}

	ctx.JSON(200, res)

}

func (c *orderController) DeleteOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.OrderService.DeleteOrder(id); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in delete order service."})
		return
	}

	ctx.JSON(200, "Deleted order.")
}
