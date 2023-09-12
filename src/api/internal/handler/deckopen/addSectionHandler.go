package deckopen

import (
	"net/http"

	"SpaceRepetition/src/api/internal/logic/deckopen"
	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddSectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SectionAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := deckopen.NewAddSectionLogic(r.Context(), svcCtx)
		resp, err := l.AddSection(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
