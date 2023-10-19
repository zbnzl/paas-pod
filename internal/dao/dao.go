package dao

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewMysqlDb, NewK8sClient)
