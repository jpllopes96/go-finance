package api

//db is the name of project then /db/sqlc
import (
	db "go-finances/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//routes
	//users
	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/id/:id", server.getUserById)

	//Categories
	router.POST("/category", server.createCategory)
	router.GET("/category/id/:id", server.getCategory)
	router.GET("/categories", server.getCategories)
	router.DELETE("/category/id/:id", server.deleteCategory)
	router.PUT("/category", server.updateCategory)

	//Accounts
	router.POST("/account", server.createAccount)
	router.GET("/account/id/:id", server.getAccount)
	router.GET("/accounts", server.getAccounts)
	router.DELETE("/account/id/:id", server.deleteAccount)
	router.PUT("/account", server.updateAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error": err.Error()}
}
