package data

import (
	"code.srdcloud.cn/AItestproject/AIPass/aicore-common/log"
	"context"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData() (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(log.StdLogger).Info(context.Background(), "closing the data resources")
	}
	return &Data{}, cleanup, nil
}
