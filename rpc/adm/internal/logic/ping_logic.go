package logic

import (
	"context"

	"algorithm/rpc/adm/adm"
	"algorithm/rpc/adm/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *adm.Request) (*adm.Response, error) {
	// todo: add your logic here and delete this line

	return &adm.Response{}, status.Error(codes.Internal, "mock fuse")
}
