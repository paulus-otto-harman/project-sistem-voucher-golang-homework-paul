package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
	"project/util"
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
	responseCreated(c, voucher, "Voucher created successfully")
}

func (ctrl *AdminVoucherController) Index(c *gin.Context) {
	page, err := util.Uint(c.Query("p"))
	limit, err := util.Uint(c.Query("l"))
	isActive := c.Query("active")
	area := c.Query("area")
	voucherType := c.Query("type")

	total, pages, currentPage, itemsPerPage, vouchers, err := ctrl.service.All(page, limit, isActive, area, voucherType)
	if err != nil {
		ctrl.logger.Error("Failed to get vouchers", zap.Error(err))
		responseError(c, "GET_ERROR", "Failed to get vouchers", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Vouchers retrieved successfully", zap.Any("vouchers", vouchers))
	responseDataPage(c, total, pages, currentPage, itemsPerPage, vouchers)
}

func (ctrl *AdminVoucherController) Update(c *gin.Context) {
	var voucher domain.Voucher
	var err error

	voucher.ID, err = util.Uint(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Invalid user ID", zap.Error(err))
		responseError(c, "INVALID_ID", "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err = c.ShouldBindJSON(&voucher); err != nil {
		ctrl.logger.Error("Failed to bind user data", zap.Error(err))
		responseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.Update(voucher); err != nil {
		ctrl.logger.Error("Failed to update user", zap.Error(err))
		responseError(c, "UPDATE_ERROR", "Failed to update user", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("User updated successfully", zap.Any("user", voucher))
	responseOK(c, voucher, "User updated successfully")
}

func (ctrl *AdminVoucherController) Delete(c *gin.Context) {
	id, err := util.Uint(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Invalid voucher ID", zap.Error(err))
		responseError(c, "INVALID_ID", "Invalid voucher ID", http.StatusBadRequest)
		return
	}

	if err := ctrl.service.Delete(id); err != nil {
		ctrl.logger.Error("Failed to delete voucher", zap.Error(err))
		responseError(c, "DELETE_ERROR", "Failed to delete voucher", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Voucher deleted successfully", zap.Uint("voucherID", id))
	responseOK(c, nil, "Voucher deleted successfully")

}
