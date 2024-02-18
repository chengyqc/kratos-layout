package main

import (
	"code.srdcloud.cn/AItestproject/AIPass/aicore-common/log"
	"code.srdcloud.cn/AItestproject/AIPass/aicore-common/panic_catch"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"os"

	gin2 "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	// logOutput
	logOutput string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&logOutput, "log_output", "file", "log output, opt: file console")
}

func newApp(hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(log.StdLogger),
		kratos.Server(
			hs,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var globalConfig = &conf.Conf{}
	if err := c.Scan(globalConfig); err != nil {
		panic(err)
	}

	// logger
	sysLogger := log.InitLogger(log.Log{
		Level:          globalConfig.Config.Log.Level,
		Formatter:      globalConfig.Config.Log.Formatter,
		CutTime:        globalConfig.Config.Log.CutTime,
		LogFileSaveNum: globalConfig.Config.Log.LogFileSaveNum,
	}, log.WithOutputMode(logOutput))

	// watch config
	go func() {
		defer panic_catch.CatchPanic()
		ctx := context.Background()
		helper := log.NewHelper(sysLogger)
		if err := c.Watch("config", func(s string, value config.Value) {
			if err := value.Scan(globalConfig.Config); err != nil {
				helper.Error(ctx, err)
			}
			helper.Debug(ctx, "config update")
		}); err != nil {
			helper.Error(ctx, err)
		}
	}()

	// gin set GIN_MODE=release
	gin.SetMode("release")
	engine := gin.New()
	// 使用kratos中间件
	engine.Use(gin2.Middlewares(
		recovery.Recovery(),
		tracing.Server(),
		//logging.Server(httpLogger),
		validate.Validator(),
	))

	app, cleanup, err := wireApp(globalConfig.Config, engine)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
