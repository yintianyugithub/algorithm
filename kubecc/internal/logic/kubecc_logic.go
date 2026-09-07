// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"algorithm/kubecc/internal/svc"
	"algorithm/kubecc/internal/types"
	"algorithm/rpc/adm/adm"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type KubeccLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKubeccLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KubeccLogic {
	return &KubeccLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KubeccLogic) Kubecc(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	rsp, err := l.svcCtx.AdmRpc.Ping(l.ctx, &adm.Request{})

	if rsp == nil {
		resp = &types.Response{
			Message: "nil",
		}
		return
	}

	return &types.Response{
		Message: rsp.Pong,
	}, nil
}
