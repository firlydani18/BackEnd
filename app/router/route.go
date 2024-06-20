package router

import (
	ah "KosKita/features/admin/handler"
	as "KosKita/features/admin/service"
	bd "KosKita/features/booking/data"
	bh "KosKita/features/booking/handler"
	bs "KosKita/features/booking/service"
	cd "KosKita/features/chat/data"
	ch "KosKita/features/chat/handler"
	cs "KosKita/features/chat/service"
	kd "KosKita/features/kos/data"
	kh "KosKita/features/kos/handler"
	ks "KosKita/features/kos/service"
	ud "KosKita/features/user/data"
	uh "KosKita/features/user/handler"
	us "KosKita/features/user/service"

	"KosKita/utils/cloudinary"
	"KosKita/utils/encrypts"
	"KosKita/utils/externalapi"
	"KosKita/utils/middlewares"
	ws "KosKita/utils/websocket"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {

	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()
	midtrans := externalapi.New()

	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, cloudinaryUploader)

	kosData := kd.New(db)
	kosService := ks.New(kosData, userService)
	kosHandlerAPI := kh.New(kosService, cloudinaryUploader)

	bookData := bd.NewBooking(db, midtrans)
	bookService := bs.NewBooking(bookData)
	bookHandlerAPI := bh.NewBooking(bookService)

	adminService := as.New(userData, kosData, bookData, userService)
	adminHandlerAPI := ah.New(adminService)

	chatData := cd.New(db)
	chatService := cs.New(chatData)
	hub := ws.NewHub()
	wsHandler := ch.New(chatService, hub, userService)
	go hub.Run()

	// define routes/ endpoint MESSAGE
	e.POST("/create-room", wsHandler.CreateRoom)
	e.GET("/get-room", wsHandler.GetRooms, middlewares.JWTMiddleware())
	e.GET("/join-room/:roomId", wsHandler.JoinRoom)
	e.GET("/room/:roomId", wsHandler.GetMessages)

	// define routes/ endpoint ADMIN
	e.GET("/admin", adminHandlerAPI.GetAllData, middlewares.JWTMiddleware())

	// define routes/ endpoint IMAGE
	e.POST("/upload-image/:kosid", kosHandlerAPI.UploadImages, middlewares.JWTMiddleware())
	e.PUT("/upload-image/:kosid", kosHandlerAPI.UpdateImages, middlewares.JWTMiddleware())

	// define routes/ endpoint USER
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/change-password", userHandlerAPI.ChangePassword, middlewares.JWTMiddleware())

	// define routes/ endpoint KOS
	e.POST("/kos", kosHandlerAPI.CreateKos, middlewares.JWTMiddleware())
	e.PUT("/kos/:id", kosHandlerAPI.UpdateKos, middlewares.JWTMiddleware())
	e.POST("/kos/:id/rating", kosHandlerAPI.CreateRating, middlewares.JWTMiddleware())
	e.GET("/kos", kosHandlerAPI.GetKosByRating)
	e.DELETE("/kos/:id", kosHandlerAPI.DeleteKos, middlewares.JWTMiddleware())
	e.GET("/kos/:id", kosHandlerAPI.GetKosById)
	e.GET("/users/kos", kosHandlerAPI.GetKosByUserId, middlewares.JWTMiddleware())
	e.GET("/kos/search", kosHandlerAPI.SearchKos)

	// define routes/ endpoint BOOKING
	e.POST("/booking", bookHandlerAPI.CreateBooking, middlewares.JWTMiddleware())
	e.GET("/booking", bookHandlerAPI.GetBookings, middlewares.JWTMiddleware())
	e.PUT("/booking/:booking_id", bookHandlerAPI.CancelBookingById, middlewares.JWTMiddleware())
	e.POST("/payment/notification/webhook", bookHandlerAPI.WebhoocksNotification)
}
