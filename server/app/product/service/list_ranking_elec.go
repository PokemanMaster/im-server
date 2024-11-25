package service

import (
	"context"
	"fmt"
	"github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"

	"strings"
)

// ListElecRankingService 展示排行的服务
type ListElecRankingService struct {
}

// List 获取家电排行
func (service *ListElecRankingService) List(ctx context.Context) resp.Response {
	var products []model.Product
	code := e.SUCCESS
	// 从redis读取点击前十的视频
	pros, _ := cache.RC.ZRevRange(ctx, cache.ElectricalRank, 0, 6).Result()

	if len(pros) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(pros, ","))
		err := db.DB.Where("id in (?)", pros).Order(order).Find(&products).Error
		if err != nil {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
			code := e.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProducts(products),
	}
}
