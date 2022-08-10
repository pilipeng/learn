package handler

import (
	"hello/internal/logic"
	"hello/internal/svc"
	"hello/internal/types"
	"hello/response" // ①
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHelloLogic(r.Context(), ctx)
		resp, err := l.Hello(&req)
		response.Response(w, resp, err) //②

	}
}
