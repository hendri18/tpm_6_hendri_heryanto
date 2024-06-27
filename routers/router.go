package routers

import (
	"tpm_6_HendriHeryanto/handler"
	"tpm_6_HendriHeryanto/repository"
	"tpm_6_HendriHeryanto/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	productRepo := &repository.ProductRepo{DB: db}
	productService := &service.ProductService{ProductRepo: productRepo}
	productHandler := &handler.ProductHandler{ProductService: productService}

	router.GET("/product", productHandler.Get)
	router.POST("/product", productHandler.Create)
	router.PUT("/product/:id", productHandler.Update)
	router.DELETE("/product/:id", productHandler.Delete)

	return router
}
