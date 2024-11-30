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
		adminApi.GET("/vouchers", ctx.Ctl.AdminVoucher.Get)
		adminApi.GET("/vouchers/usages", ctx.Ctl.AdminVoucherUsage.Get)
	}

	userApi := r.Group("/user")
	{
		userApi.POST("/redemptions", nil)
		userApi.GET("/redemptions", nil)

		userApi.GET("/vouchers", nil)
		userApi.GET("/vouchers/:id", nil)
		userApi.POST("/vouchers/:id/usages", nil)
		userApi.GET("/vouchers/usages", nil)
	}

	return r
}
