package knowledge

import (
	"net/http"

	"SpaceRepetition/src/api/internal/logic/knowledge"
	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddknowledgeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.KnowledgeAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := knowledge.NewAddknowledgeLogic(r.Context(), svcCtx)
		resp, err := l.Addknowledge(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
