package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
)

type Handler struct {
	AdminVoucher      AdminVoucherController
	AdminVoucherUsage AdminVoucherUsageController
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		AdminVoucher:      *NewAdminVoucherController(service.Voucher, logger),
		AdminVoucherUsage: *NewAdminVoucherUsageController(service.Voucher, logger),
	}
}

func responseOK(c *gin.Context, data interface{}, description string) {
	c.JSON(http.StatusOK, domain.HTTPResponse{
		Status:      true,
		Description: description,
		Data:        data,
	})
}

func responseError(c *gin.Context, errorCode string, description string, httpStatusCode int) {
	c.JSON(httpStatusCode, domain.HTTPResponse{
		Status:      false,
		ErrorCode:   errorCode,
		Description: description,
	})
}
