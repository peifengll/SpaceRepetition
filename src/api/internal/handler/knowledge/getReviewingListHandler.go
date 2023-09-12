package knowledge

import (
	"net/http"

	"SpaceRepetition/src/api/internal/logic/knowledge"
	"SpaceRepetition/src/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetReviewingListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := knowledge.NewGetReviewingListLogic(r.Context(), svcCtx)
		resp, err := l.GetReviewingList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
