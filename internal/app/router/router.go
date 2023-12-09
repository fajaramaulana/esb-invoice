package router

import (
	"esb-invoice/internal/app/controller"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type router struct {
	router   *gin.Engine
	customer *controller.CustomerController
}

func NewRouter(customer *controller.CustomerController) *router {
	return &router{
		router:   gin.Default(),
		customer: customer,
	}
}

func (r *router) SetupRouter(port string) {
	// docs.SwaggerInfo.BasePath = ""
	v1 := r.router.Group("/api/v1")

	// health check
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Customer
	v1.POST("/customer", r.customer.CreateCustomer)
	v1.PUT("/customer/:id", r.customer.UpdateCustomer)
	v1.DELETE("/customer/:id", r.customer.DeleteCustomer)
	v1.GET("/customer/:id", r.customer.FindCustomerById)
	v1.GET("/customer", r.customer.FindAllCustomer)

	r.router.GET("/swaggerr/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.router.Run(port)
}
