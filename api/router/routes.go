package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaingotnochill/tempmee/api/middleware"
	"github.com/swaingotnochill/tempmee/bootstrap"
	"go.mongodb.org/mongo-driver/mongo"
)

// API routes
// 1. User should be able to create account. /createAccount
// 2. User should be able to see list of books. /getAllBooks
// 3. User should be able to make order. /orderBook
// 4. User should be able to see order history. /showOrderHistory

func SetUp(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	// List of all Public APIs.
	NewHelloRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewSignupRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")

	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// List of all Private APIs.
	NewBookRouter(env, timeout, db, protectedRouter)
	NewOrderRouter(env, timeout, db, protectedRouter)
}
