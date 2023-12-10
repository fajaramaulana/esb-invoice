package main

import (
	"esb-invoice/internal/app/config/db"
	"esb-invoice/internal/app/controller"
	"esb-invoice/internal/app/router"
	"esb-invoice/internal/domain/repository/gorm"
	"esb-invoice/internal/domain/service"
	"esb-invoice/seeders"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title ESB-Invoice
// @version 1.0
// @description ESB - Invoice
// @termsOfService http://swagger.io/terms/
// @contact.name Swagger API Team
// @contact.url http://swagger.io
// @contact.email fajaragusmaulana@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081

func main() {
	// Load environment variables
	checkGoEnv := os.Getenv("GO_ENV")
	if checkGoEnv == "" { // if GO_ENV is empty, set to development
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	// connect to database
	dbInstance, err := db.ConnectionMysqlGorm()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = seeders.SeederDB(dbInstance)

	if err != nil {
		log.Fatal("Error Seeding DB:", err)
	}

	// repository
	customerRepository := gorm.NewCustomerRepository(dbInstance)
	typeItemRepository := gorm.NewTypeRepository(dbInstance)
	itemRepository := gorm.NewItemRepository(dbInstance)

	// service
	customerService := service.NewCustomerService(customerRepository)
	typeItemService := service.NewTypeService(typeItemRepository)
	itemService := service.NewItemService(itemRepository, typeItemRepository)

	// controller
	customerController := controller.NewCustomerController(customerService)
	typeItemController := controller.NewTypeController(typeItemService)
	itemController := controller.NewItemController(itemService)

	// router
	app := router.NewRouter(customerController, typeItemController, itemController)

	app.SetupRouter(os.Getenv("PORT"))
}
