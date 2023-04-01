package controller

import (
	"fmt"
	"net/http"
	"qkeruen/dto"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type offerDriverController struct {
	OfferService service.OfferDriverSevvice
}

func NewOfferDriverController(offer service.OfferDriverSevvice) offerDriverController {
	return offerDriverController{
		OfferService: offer,
	}
}

func (c *offerDriverController) CreateOffer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var create dto.OfferRequest
	if err := ctx.ShouldBindJSON(&create); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}

	err := c.OfferService.CreateOffer(id, create)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in offer create service."})
		return
	}

	ctx.JSON(201, gin.H{"message": "Saved."})

}

func (c *offerDriverController) GetMyOffer(ctx *gin.Context) {
	driverId, _ := strconv.Atoi(ctx.Param("id"))

	data, err := c.OfferService.MyOffer(driverId)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in my offer service."})
		return
	}

	ctx.JSON(201, data)
}

func (c *offerDriverController) AllOffer(ctx *gin.Context) {
	allOffer, err := c.OfferService.FindAllOffers()

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in Get all offer service."})
		return
	}

	ctx.JSON(200, allOffer)
}

func (c *offerDriverController) SearchOffers(ctx *gin.Context) {
	var offer dto.OfferRequest

	if err := ctx.ShouldBindJSON(&offer); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}

	res, err := c.OfferService.SearchOffers(offer.To, offer.From)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in search offer service."})
		return
	}

	ctx.JSON(200, res)
}

func (c *offerDriverController) DeleteOffer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.OfferService.Delete(id); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in delete offer service."})
		return
	}

	ctx.JSON(200, gin.H{"message": "Deleted."})

}
