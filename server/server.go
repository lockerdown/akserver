package server

import (
	_ "akserver/docs"
	"akserver/server/akcenter/monitoring"
	"akserver/setting"
	"crypto/tls"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
)

// A AkCenter defines parameters for running an HTTP server.
// The zero value for Server is a valid configuration.
type AkCenter struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Addr string

	TLSConfig *tls.Config
}

func (a *AkCenter) Start() error {
	serverAdmin := &http.Server{
		Addr:    a.Addr,
		Handler: a.handler(),
	}
	err := serverAdmin.ListenAndServeTLS(setting.TlsPublicKey, setting.TlsPrivateKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func (a *AkCenter) Stop() error {
	return nil
}

func (a *AkCenter) handler() http.Handler {
	gin.DisableConsoleColor()

	gin.SetMode(setting.ServerMode)
	r := gin.New()
	//r.HTMLRender = createMyRender()
	//r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return ""
	}))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404!",
		})
	})
	//r.Use(gin.Recovery())
	LoadUrl(r)
	return r
}

func LoadUrl(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})
	monitoring.Router(r)
	//login.Router(r)
}

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "static/html/index.html")
	return p
}

func Run() {
	serverAdmin := &AkCenter{
		Addr: setting.AkServerAddr + ":" + setting.AkServerPort,
	}

	serverAdmin.Start()
}
