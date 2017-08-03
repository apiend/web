package main

import (
	"fmt"
	// "log"
	"net/http"

	. "apiServer/conf"
	"apiServer/module/log"

	c "github.com/dchest/captcha"

	"github.com/hb-go/echo-mw/captcha"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	confFilePath := "" // 默认conf/conf.toml

	// 初始化配置
	if err := InitConfig(confFilePath); err != nil {
		// log.Panic(err)
		log.Debugf("run with conf:%s", confFilePath)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 验证码，优先于静态资源
	e.Use(captcha.Captcha(captcha.Config{
		CaptchaPath: "/captcha/",
		SkipLogging: false,
	}))

	fmt.Println(c.NewLen(6))

	// Route => handler
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, Conf.App.Name)
	})
	// log.Global.Printf("ddd")
	// Start server
	e.Logger.Fatal(e.Start(Conf.Server.Addr))
}
