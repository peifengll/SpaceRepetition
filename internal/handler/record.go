package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/_grpc/train_client"
	"github.com/peifengll/SpaceRepetition/internal/export_task"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"log"
	"net/http"
	"strconv"
	"time"
)

type RecordHandler struct {
	*Handler
	recordService service.RecordService
}

func NewRecordHandler(handler *Handler, recordService service.RecordService) *RecordHandler {
	return &RecordHandler{
		Handler:       handler,
		recordService: recordService,
	}
}

func (h *RecordHandler) GetRecord(ctx *gin.Context) {
}

func (h *RecordHandler) ExportReviewInfo(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var req v1.ExportRevlogCsvReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if req.Date2 < req.Date1 {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, "开始时间不能大于结束时间")
		return
	}
	req.Start = time.Unix(req.Date1/1000, 0)
	req.End = time.Unix(req.Date2/1000, 0)

	//flag, err := h.recordService.ExportInfoRecord(&req, userId)
	tsk := v1.ExportTask{
		Start:  req.Start,
		End:    req.End,
		UserId: userId,
	}
	ok, err := etask.TM.AddTask(tsk)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ExportReviewInfoErr, err.Error())
		return
	}
	if ok {
		v1.HandleSuccess(ctx, "加入到导出任务队列成功")
		return
	} else {
		v1.HandleSuccess(ctx, "当前排队人多，请换个时间再来噢！")
		return
	}
}

func (h *RecordHandler) GetExportInfos(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	// 根据userid 查询出exportinfos
	infos, err := h.recordService.GetExportInfos(userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, infos)
}

func (h *RecordHandler) GetRevlogCsvFile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var req v1.ExportFileDownloadReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	ctx.File(req.FilePath)
	//if err != nil {
	//	v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
	//	return
	//}
	//v1.HandleSuccess(ctx, nil)
}

func (h *RecordHandler) Train(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	pa := ctx.Query("filepath")
	epids := ctx.Query("epid")
	if epids == "" || pa == "" {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	epid, err := strconv.Atoi(epids)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, err)
		return
	}
	ok, err := train_client.TrainDataRequest(pa)
	// 把pa的.csv 换掉，换成_csv ，然后进行压缩 ，保留位置
	log.Println(ok)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, err)
		return
	}
	err = h.recordService.AfterTrainData(pa, epid)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
