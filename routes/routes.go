package routes

import (
	"project/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	adminApi := r.Group("/admin")
	{
		adminApi.POST("/vouchers", ctx.Ctl.AdminVoucher.Create)
		adminApi.DELETE("/vouchers/:id", ctx.Ctl.AdminVoucher.Delete)
		adminApi.PUT("/vouchers/:id", ctx.Ctl.AdminVoucher.Update)
		adminApi.GET("/vouchers", ctx.Ctl.AdminVoucher.Index)
		adminApi.GET("/vouchers/customer", ctx.Ctl.AdminVoucherCustomer.Get)
	}

	userApi := r.Group("/user")
	{
		userApi.POST("/redemptions", ctx.Ctl.CustomerRedeemVoucher.Create)
		userApi.GET("/redemptions", nil)

		userApi.POST("/vouchers", nil)
		userApi.GET("/vouchers", nil)
	}

	r.GET("/vouchers", ctx.Ctl.Voucher.Index)
	r.GET("/vouchers/:id", nil)

	return r
}
