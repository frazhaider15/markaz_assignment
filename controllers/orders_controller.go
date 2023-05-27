package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markaz/dto"
	"github.com/markaz/services"
)

func CreateOrder(ctx *gin.Context) {
	var request dto.OrderRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ctx.JSON(http.StatusOK, NewStandardResponse(false, 403, err.Error(), nil))
		return
	}
	err = request.Validate()
	if err != nil {
		ctx.JSON(http.StatusOK, NewStandardResponse(false, 403, err.Error(), nil))
		return
	}
	err = services.CreateOrder(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewStandardResponse(false, 500, err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, NewStandardResponse(true, http.StatusOK, "Order created successfully", nil))
}

func GetCustomerOrders(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	orders, err := services.GetCustomerOrders(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewStandardResponse(false, 500, err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, NewStandardResponse(true, http.StatusOK, "Orders fetched successfully", orders))
}
