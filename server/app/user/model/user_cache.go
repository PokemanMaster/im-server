package model

import (
	"context"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"

	"time"
)

func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	cache.RC.Set(ctx, key, val, timeTTL)
}
