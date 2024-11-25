package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
)

type DeleteFriendService struct {
	UserID   uint
	FriendID uint
}

func (service *DeleteFriendService) Delete() *resp.Response {
	// 检查好友关系是否存在
	var contact model.Contact
	err := db.DB.Model(&model.Contact{}).
		Where("owner_id = ? AND target_id = ? AND type = ?", service.UserID, service.FriendID, 1).
		First(&contact).Error
	if err != nil {
		zap.L().Info("好友不存在", zap.String("app.friend.service.delete_friend", err.Error()))
		return &resp.Response{
			Status: e.ERROR_NOT_EXIST_FRIEND,
			Msg:    e.GetMsg(e.ERROR_NOT_EXIST_FRIEND),
		}
	}

	// 删除好友关系
	tx := db.DB.Begin()

	err = tx.Delete(&contact).Error
	if err != nil {
		tx.Rollback()
		zap.L().Info("删除好友失败", zap.String("app.friend.service.delete_friend", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 双向删除好友
	err = tx.Where("owner_id = ? AND target_id = ? AND type = ?", service.FriendID, service.UserID, 1).
		Delete(&model.Contact{}).Error
	if err != nil {
		tx.Rollback()
		zap.L().Info("双向删除好友失败", zap.String("app.friend.service.delete_friend", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	tx.Commit()

	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
