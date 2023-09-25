package api

import (
	_ "warehouse/api/docs"
	"warehouse/api/handler"
	"warehouse/config"
	"warehouse/pkg/logger"
	"warehouse/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, storage, logger)

	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	r.POST("/branch", handler.CreateBranch)
	r.GET("/branch/:id", handler.GetByIdBranch)
	r.GET("/branch", handler.GetListBranch)
	r.PUT("/branch/:id", handler.UpdateBranch)
	r.DELETE("/branch/:id", handler.DeleteBranch)

	r.POST("/coming", handler.CreateComing)
	r.GET("/coming/:id", handler.GetByIdComing)
	r.GET("/coming", handler.GetListComing)
	r.PUT("/coming/:id", handler.UpdateComing)
	r.DELETE("/coming/:id", handler.DeleteComing)

	r.POST("/comingproducts", handler.CreateComingProducts)
	r.GET("/comingproducts/:id", handler.GetByIdComingProducts)
	r.GET("/comingproducts", handler.GetListComingProducts)
	r.PUT("/comingproducts/:id", handler.UpdateComingProducts)
	r.DELETE("/comingproducts/:id", handler.DeleteComingProducts)

	r.POST("/remainder", handler.CreateRemainder)
	r.GET("/remainder/:id", handler.GetByIdRemainder)
	r.GET("/remainder", handler.GetListRemainder)
	r.PUT("/remainder/:id", handler.UpdateRemainder)
	r.DELETE("/remainder/:id", handler.DeleteRemainder)

	r.POST("/do-income/:coming_id", handler.Income_Warehouse)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
