package logic

import (
	"context"
	"fmt"
	"go-mall/app/user/cmd/api/internal/svc"
	"go-mall/app/user/cmd/api/internal/types"
        "go-mall/app/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserResponse, err error) {
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
            Id: "1",
        })

	fmt.Println("user ->", user)
        return &types.UserResponse{
            Id:   req.Id,
            Name: "test user",
        }, nil
}
