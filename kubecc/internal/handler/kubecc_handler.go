// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package handler

import (
	"net/http"

	"algorithm/kubecc/internal/logic"
	"algorithm/kubecc/internal/svc"
	"algorithm/kubecc/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func KubeccHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewKubeccLogic(r.Context(), svcCtx)
		resp, err := l.Kubecc(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
