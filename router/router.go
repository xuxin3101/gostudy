package router

import ("github.com/gin-gonic/gin"
	"gostudy/handler/sd"
	"gostudy/router/middleware"
	"net/http"
)

func Load(g *gin.Engine,mw ...gin.HandlerFunc) *gin.Engine{
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound,"错误的路由")
	})
	svcd :=g.Group("/sd")
	{
		svcd.GET("/heathe",sd.HealthCheck)
		svcd.GET("/disk",sd.DiskCheck)
		svcd.GET("/cpu",sd.CPUCheck)
		svcd.GET("/ram",sd.RAMCheck)

	}
	return g

}