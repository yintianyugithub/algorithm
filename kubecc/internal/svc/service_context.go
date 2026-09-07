// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"algorithm/kubecc/internal/config"
	"algorithm/rpc/adm/adm_client"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	// RPC客户端
	AdmRpc adm_client.Adm
}

func NewServiceContext(c config.Config) *ServiceContext {
	adm := adm_client.NewAdm(zrpc.MustNewClient(c.AdmRpc))

	return &ServiceContext{
		Config: c,
		AdmRpc: adm,
	}
}
