package svc

import (
        "go-mall/app/order/cmd/api/internal/config"
        "go-mall/app/user/cmd/rpc/userclient"

        "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
        Config config.Config
        UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
        return &ServiceContext{
                Config: c,
                UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
        }
}
