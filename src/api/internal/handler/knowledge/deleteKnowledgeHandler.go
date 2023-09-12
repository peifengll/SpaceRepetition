package knowledge

import (
	"net/http"

	"SpaceRepetition/src/api/internal/logic/knowledge"
	"SpaceRepetition/src/api/internal/svc"
	"SpaceRepetition/src/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteKnowledgeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.KnowledgeDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := knowledge.NewDeleteKnowledgeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteKnowledge(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
