package api

//db is the name of project then /db/sqlc
import (
	db "go-finances/db/sqlc"
	"go-finances/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func CORSConfig() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}

		context.Next()
	}
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//cors//
	router.Use(CORSConfig())
	//routes
	//login
	router.POST("/login", server.login)

	//users
	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/id/:id", server.getUserById)
	router.GET("/user/email/:email", server.getUsetByEmail)

	//Categories
	router.POST("/category", middleware.GetTokenHeaderAndValidate, server.createCategory)
	router.GET("/category/id/:id", middleware.GetTokenHeaderAndValidate, server.getCategory)
	router.GET("/categories", middleware.GetTokenHeaderAndValidate, server.getCategories)
	router.DELETE("/category/id/:id", middleware.GetTokenHeaderAndValidate, server.deleteCategory)
	router.PUT("/category", middleware.GetTokenHeaderAndValidate, server.updateCategory)

	//Accounts
	router.POST("/account", middleware.GetTokenHeaderAndValidate, server.createAccount)
	router.GET("/account/id/:id", middleware.GetTokenHeaderAndValidate, server.getAccount)
	router.GET("/account/graph/:user_id/:type", middleware.GetTokenHeaderAndValidate, server.getAccountGraph)
	router.GET("/account/reports/:user_id/:type", middleware.GetTokenHeaderAndValidate, server.getAccountReport)
	router.GET("/accounts", middleware.GetTokenHeaderAndValidate, server.getAccounts)
	router.DELETE("/account/id/:id", middleware.GetTokenHeaderAndValidate, server.deleteAccount)
	router.PUT("/account", middleware.GetTokenHeaderAndValidate, server.updateAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error": err.Error()}
}
