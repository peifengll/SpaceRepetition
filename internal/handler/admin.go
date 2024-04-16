package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	*Handler
	adminService service.AdminService
}

func NewAdminHandler(
	handler *Handler,
	adminService service.AdminService,
) *AdminHandler {
	return &AdminHandler{
		Handler:      handler,
		adminService: adminService,
	}
}

func (h *AdminHandler) GetAdmin(ctx *gin.Context) {

}

func (h *AdminHandler) Register(ctx *gin.Context) {
	req := new(v1.AdminRegisterReq)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	id, err := strconv.Atoi(GetUserIdFromCtx(ctx))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	if err := h.adminService.Register(ctx, req, id); err != nil {
		h.logger.WithContext(ctx).Error("adminService.Register error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

func (h *AdminHandler) Login(ctx *gin.Context) {
	var req v1.LoginAdminReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	token, err := h.adminService.Login(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	v1.HandleSuccess(ctx, v1.LoginResponseData{
		AccessToken: token,
	})
}

func (h *AdminHandler) UpdatePrivileges(ctx *gin.Context) {
	var req v1.UpdatePrivilegesReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	id, err := strconv.Atoi(GetUserIdFromCtx(ctx))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	err = h.adminService.UpdatePrivileges(ctx, &req, id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrUnauthorized, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *AdminHandler) DelAdminAccount(ctx *gin.Context) {
	var req v1.DelAdminReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	id, err := strconv.Atoi(GetUserIdFromCtx(ctx))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	err = h.adminService.DelAdminAccount(ctx, &req, id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrUnauthorized, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *AdminHandler) ShowAdminAccounts(ctx *gin.Context) {
	res, err := h.adminService.GetAccounts(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrUnauthorized, nil)
		return
	}
	v1.HandleSuccess(ctx, res)
}
