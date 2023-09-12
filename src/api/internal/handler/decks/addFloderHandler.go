package decks

import (
	"net/http"

	"SpaceRepetition/src/api/internal/logic/decks"
	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddFloderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FloderAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := decks.NewAddFloderLogic(r.Context(), svcCtx)
		resp, err := l.AddFloder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
