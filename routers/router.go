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

	router.GET("/products", productHandler.Get)
	router.POST("/products", productHandler.Create)
	router.PUT("/products/:id", productHandler.Update)
	router.DELETE("/products/:id", productHandler.Delete)

	return router
}
