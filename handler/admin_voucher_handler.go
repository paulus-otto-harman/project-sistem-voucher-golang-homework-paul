package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
)

type AdminVoucherController struct {
	service service.VoucherService
	logger  *zap.Logger
}

func NewAdminVoucherController(service service.VoucherService, logger *zap.Logger) *AdminVoucherController {
	return &AdminVoucherController{service, logger}
}

func (ctrl *AdminVoucherController) Create(c *gin.Context) {
	var voucher domain.Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		ctrl.logger.Error("Failed to bind user data", zap.Error(err))
		responseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.Create(voucher); err != nil {
		ctrl.logger.Error("Failed to create user", zap.Error(err))
		responseError(c, "CREATE_ERROR", "Failed to create voucher", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("User created successfully", zap.Any("voucher", voucher))
	responseOK(c, voucher, "Voucher created successfully")
}

func (ctrl *AdminVoucherController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of vouchers"})
}

func (ctrl *AdminVoucherController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update voucher"})
}

func (ctrl *AdminVoucherController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete voucher"})
}
