package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
)

type CustomerRedeemVoucherController struct {
	service service.RedeemVoucherService
	logger  *zap.Logger
}

func NewCustomerRedeemVoucherController(service service.RedeemVoucherService, logger *zap.Logger) *CustomerRedeemVoucherController {
	return &CustomerRedeemVoucherController{service, logger}
}

func (ctrl *CustomerRedeemVoucherController) Create(c *gin.Context) {
	var voucher domain.Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		ctrl.logger.Error("Failed to bind voucher data", zap.Error(err))
		responseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.Create(voucher); err != nil {
		ctrl.logger.Error("Failed to redeem voucher", zap.Error(err))
		responseError(c, "CREATE_ERROR", "Failed to redeem voucher", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Voucher redeemed successfully", zap.Any("voucher", voucher))
	responseCreated(c, voucher, "Voucher redeemed successfully")
}
