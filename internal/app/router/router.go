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
	typeItem *controller.TypeController
	item     *controller.ItemController
	invoice  *controller.InvoiceController
}

func NewRouter(customer *controller.CustomerController, typeItem *controller.TypeController, item *controller.ItemController, invoice *controller.InvoiceController) *router {
	return &router{
		router:   gin.Default(),
		customer: customer,
		typeItem: typeItem,
		item:     item,
		invoice:  invoice,
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

	// Type
	v1.GET("/type", r.typeItem.FindAll)
	v1.POST("/type", r.typeItem.Create)
	v1.PUT("/type/:id", r.typeItem.Update)
	v1.DELETE("/type/:id", r.typeItem.Delete)

	// Item
	v1.GET("/item", r.item.FindAllItem)
	v1.GET("/item/:id", r.item.FindByIdItem)
	v1.POST("/item", r.item.CreateItem)
	v1.PUT("/item/:id", r.item.UpdateItem)
	v1.DELETE("/item/:id", r.item.DeleteItem)

	v1.GET("/invoice", r.invoice.FindAllInvoice)
	v1.GET("/invoice/:id", r.invoice.FindInvoiceById)
	v1.POST("/invoice", r.invoice.CreateInvoice)
	v1.PUT("/invoice/:id", r.invoice.UpdateInvoice)

	r.router.GET("/swaggerr/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.router.Run(port)
}
