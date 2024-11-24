package service

import (
	"github.com/PokemanMaster/GoChat/app/product/model"
	"github.com/PokemanMaster/GoChat/app/product/serializer"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/mid"
	"github.com/PokemanMaster/GoChat/resp"
	"go.uber.org/zap"
)

// SearchProductsService 搜索商品的服务
type SearchProductsService struct {
	Search string
}

// Show 搜索商品
func (service *SearchProductsService) Show() resp.Response {

	validSearch, code, err := mid.ValidateSearchInput(service.Search)
	if code != e.SUCCESS {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productParam, code, err := model.SearchProductParam(validSearch)
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProductParams(productParam),
	}
}