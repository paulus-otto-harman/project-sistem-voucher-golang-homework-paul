package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{"message": "Create voucher"})
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
