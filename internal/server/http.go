package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos/v2/transport/http"
	"time"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Config, router *gin.Engine) *http.Server {
	httpSrv := http.NewServer(
		http.Address(fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)),
		http.Timeout(30*time.Second),
	)
	httpSrv.HandlePrefix("/", router)
	return httpSrv
}
