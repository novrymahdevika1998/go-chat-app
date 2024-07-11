package routes

import (
	"chat-app/handlers"
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/users", func(c *gin.Context) { handlers.CreateUser(c, db) })
	r.POST("/channels", func(c *gin.Context) { handlers.CreateChannel(c, db) })
	r.POST("/messages", func(c *gin.Context) { handlers.CreateMessage(c, db) })

	r.GET("/channels", func(c *gin.Context) { handlers.ListChannels(c, db) })
	r.GET("/messages", func(c *gin.Context) { handlers.ListMessages(c, db) })

	r.POST("/login", func(c *gin.Context) { handlers.Login(c, db) })

	r.StaticFile("/", "chat-ui/build/index.html")
	r.StaticFS("/static", http.Dir("chat-ui/build/static"))

	return r
}
