package main

import (
	"log"
	"qkeruen/config"
	"qkeruen/controller"
	"qkeruen/middleware"
	"qkeruen/models"
	"qkeruen/repository"
	"qkeruen/service"

	"github.com/gin-gonic/gin"
)

var (
	AppSettings models.Settings
)

var dbPool, err = config.NewDBPool(config.DataBaseConfig{
	Username: "postgres",
	Password: "123456",
	Hostname: "localhost",
	Port:     "5432",
	DBName:   "postgres",
})

var (
	authDB                = repository.NewDatabase(dbPool)
	driverDB              = repository.NewDriverRepository(dbPool)
	userDB                = repository.NewUserRepository(dbPool)
	offerDriverDB         = repository.NewOfferDriverRepository(dbPool)
	offerUserDB           = repository.NewOfferUserRepository(dbPool)
	orderDB               = repository.NewOrderRepository(dbPool)
	processDb             = repository.NewProcessRepository(dbPool)
	searchDb              = repository.NewSearchRepository(dbPool)
	jwtService            = service.NewJWTService()
	authService           = service.NewAuthService(authDB)
	driverService         = service.NewDriverService(driverDB)
	userService           = service.NewUserService(userDB)
	offerDriverService    = service.NewOfferDriverService(offerDriverDB)
	offerUserService      = service.NewOfferuserService(offerUserDB)
	orderService          = service.NewOrderService(orderDB)
	processService        = service.NewProcessService(processDb)
	searchService         = service.NewSearchService(searchDb)
	authController        = controller.NewAuthController(authService, jwtService)
	driverController      = controller.NewDriverController(driverService, jwtService)
	userController        = controller.NewUserController(userService, jwtService)
	offerDriverController = controller.NewOfferDriverController(offerDriverService)
	offerUserController   = controller.NewOfferUserController(offerUserService)
	orderController       = controller.NewOrderController(orderService)
	processController     = controller.NewProcessController(processService)
	searchController      = controller.NewSearchController(searchService)
)

func main() {
	defer dbPool.Close()

	r := gin.Default()
	r.Use(gin.Recovery())
	e := config.InitTabeles(dbPool)

	if e != nil {
		log.Println(e)
	} else {
		log.Println("Success init.")
	}

	r.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello")
	})

	authRouter := r.Group("/authorization")
	{
		authRouter.POST("/sign", authController.Register)
		authRouter.POST("/check", authController.ValidatorSMS)
		authRouter.POST("/resend", authController.ResendCode)
		authRouter.POST("/check-token", authController.CheckToken)
	}

	driverRouter := r.Group("/driver", middleware.AuthorizeJWTDriver(jwtService))
	{
		driverRouter.POST("/", driverController.Register)
		driverRouter.GET("/", driverController.GetProfile)
		driverRouter.PUT("/", driverController.Update)
		driverRouter.DELETE("/:id", driverController.Delete)
	}

	userRouter := r.Group("/user", middleware.AuthorizeJWTUser(jwtService))
	{
		userRouter.POST("/", userController.Register)
		userRouter.GET("/", userController.GetProfile)
		userRouter.PUT("/", userController.Update)
		userRouter.DELETE("/:id", userController.Delete)
	}

	historyRouter := r.Group("/history")
	{
		historyRouter.GET("/driver", middleware.AuthorizeJWTDriver(jwtService))
		historyRouter.GET("/user", middleware.AuthorizeJWTUser(jwtService))
	}

	offer_driver := r.Group("/offer/driver", middleware.AuthorizeJWTDriver(jwtService))
	{
		offer_driver.POST("/:id", offerDriverController.CreateOffer)
		offer_driver.GET("/my", offerDriverController.GetMyOffer)
		offer_driver.GET("/all", offerDriverController.AllOffer)
		offer_driver.POST("/search", offerDriverController.SearchOffers)
		offer_driver.DELETE("/:id", offerDriverController.DeleteOffer)
	}

	offer_user := r.Group("/offer/user", middleware.AuthorizeJWTUser(jwtService))
	{
		offer_user.POST("/:id", offerUserController.CreateOffer)
		offer_user.GET("/my", offerUserController.GetMyOffer)
		offer_user.GET("/all", offerUserController.AllOffer)
		offer_user.POST("/search", offerUserController.SearchOffers)
		offer_user.DELETE("/:id", offerUserController.DeleteOffer)
	}

	orderDriverRouter := r.Group("/order/user", middleware.AuthorizeJWTDriver(jwtService))
	{
		orderDriverRouter.GET("/for-driver/:id", orderController.GetOrders)
		//orderDriverRouter.GET("/")
		//orderDriverRouter.DELETE("/:orderId")
	}

	orderUserRouter := r.Group("/order/user", middleware.AuthorizeJWTUser(jwtService))
	{
		orderUserRouter.POST("/", orderController.CreateOrder)
		orderUserRouter.GET("/:id", orderController.GetMyOrders)
		orderUserRouter.DELETE("/:id", orderController.DeleteOrder)
	}

	processRouter := r.Group("/process")
	{
		processRouter.POST("/start/:driverId/:orderId", processController.AcceptOrder)
		processRouter.POST("/cancel/:orderId", processController.CancellOrder)
		processRouter.POST("/finish/:driverId/:orderId", processController.FinishOrder)
		processRouter.POST("/driver/:driverId", processController.GetOrdersInProcessDriver)
		processRouter.POST("/user/userId", processController.GetOrdersInProcessUser)
	}

	search := r.Group("/search")
	{
		search.POST("/check/:place", searchController.Check)
		search.POST("/create/:place", searchController.Create)
		search.POST("/checkGeo/:place", searchController.CheckGeo)
		search.POST("/createGeo/:place", searchController.CreateGeo)
	}

	r.Run(":8080")
}
