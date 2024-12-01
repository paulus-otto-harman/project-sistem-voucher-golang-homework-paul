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
	AdminVoucherUser  AdminVoucherUserController
	UserRedeemVoucher UserRedeemVoucherController
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		AdminVoucher:      *NewAdminVoucherController(service.Voucher, logger),
		AdminVoucherUser:  *NewAdminVoucherUserController(service.Voucher, logger),
		UserRedeemVoucher: *NewUserRedeemVoucherController(service.RedeemVoucher, logger),
	}
}

func responseOK(c *gin.Context, data interface{}, description string) {
	c.JSON(http.StatusOK, domain.HTTPResponse{
		Status:      true,
		Description: description,
		Data:        data,
	})
}

func responseCreated(c *gin.Context, data interface{}, description string) {
	c.JSON(http.StatusCreated, domain.HTTPResponse{
		Status:      true,
		Description: description,
		Data:        data,
	})
}

func responseDataPage(c *gin.Context, total int64, pages int, page uint, limit uint, data interface{}) {
	c.JSON(http.StatusOK, domain.DataPage{
		Status:      true,
		Total:       total,
		Pages:       pages,
		CurrentPage: page,
		Limit:       limit,
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
