package api

import (
	"new_project/api/docs"
	"new_project/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func NewServer(h *handler.Handler) *gin.Engine {

	r := gin.Default()
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "First API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.Use()
	r.POST("/branch", h.CreateBranch)
	r.GET("/branch", h.GetAllBranch)
	r.GET("/branch/:id", h.GetBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/category", h.CreateCategory)
	r.GET("/category", h.GetAllCategory)
	r.GET("/category/:id", h.GetCategory)
	r.PUT("/category/:id", h.UpdateCategory)
	r.DELETE("/category/:id", h.DeleteCategory)

	r.POST("/comingTable", h.CreateComingTable)
	r.GET("/comingTable", h.GetAllComingTable)
	r.GET("/comingTable/:id", h.GetComingTable)
	r.PUT("/comingTable/:id", h.UpdateComingTable)
	r.DELETE("/comingTable/:id", h.DeleteComingTable)

	r.POST("/comingTableProduct", h.CreateComingTableProduct)
	r.GET("/comingTableProduct", h.GetAllComingTableProduct)
	r.GET("/comingTableProduct/:id", h.GetComingTableProduct)
	r.PUT("/comingTableProduct/:id", h.UpdateComingTableProduct)
	r.DELETE("/comingTableProduct/:id", h.DeleteComingTableProduct)

	r.POST("/product", h.CreateProduct)
	r.GET("/product", h.GetAllProduct)
	r.GET("/product/:id", h.GetProduct)
	r.PUT("/product/:id", h.UpdateProduct)
	r.DELETE("/product/:id", h.DeleteProduct)

	r.POST("/remaining", h.CreateRemaining)
	r.GET("/remaining", h.GetAllRemaining)
	r.GET("/remaining/:id", h.GetRemaining)
	r.PUT("/remaining/:id", h.UpdateRemaining)
	r.DELETE("/remaining/:id", h.DeleteRemaining)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
