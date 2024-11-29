package routes

import (
	"project/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	adminApi := r.Group("/admin")
	{
		adminApi.POST("/vouchers", nil)
		adminApi.DELETE("/vouchers/:id", nil)
		adminApi.PUT("/vouchers/:id", nil)
		adminApi.GET("/vouchers", nil)
		adminApi.GET("/vouchers/usages", nil)
	}

	userApi := r.Group("/user")
	{
		userApi.POST("/redemptions", nil)

		userApi.GET("/vouchers", nil)
		userApi.GET("/vouchers/:id", nil)
		userApi.POST("/vouchers/:id/usages", nil)
		userApi.GET("/vouchers/usages", nil)

		userApi.GET("/redemptions", nil)
	}

	return r
}
