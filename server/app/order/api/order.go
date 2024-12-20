package api

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/order/service"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateOrder 用户创建订单接口
func CreateOrder(ctx *gin.Context) {
	services := service.CreateOrderService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.order.api.order", err.Error()))
	} else {
		res := services.Create(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// ListOrders 用户订单列表
func ListOrders(ctx *gin.Context) {
	services := service.ListOrdersService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.order.api.order", err.Error()))
	} else {
		res := services.List(ctx, ctx.Param("id"))
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// ShowOrder 展示用户订单接口
func ShowOrder(ctx *gin.Context) {
	services := service.ShowOrderService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.order.api.order", err.Error()))
	} else {
		res := services.Show(ctx.Param("num"))
		ctx.JSON(200, res) // 解析数据JSON
	}
}
