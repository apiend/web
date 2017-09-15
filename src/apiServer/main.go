package main

import (

	// "log"

	"fmt"
	"net/http"

	. "apiServer/conf"
	"apiServer/module/log"
	"apiServer/util"

	"github.com/echo-contrib/pongor"
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
		SkipLogging: true,
	}))

	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("public/views/*.html")),
	// }

	r := pongor.GetRenderer(pongor.PongorOption{
		// 默认模版目录
		Directory: "public/templates",
		// if you want to reload template every request, set Reload to true.
		Reload: true,
	})
	// fmt.Println(c.NewLen(6))
	e.Renderer = r

	// 静态资源
	e.Static("/assets", "public/assets")

	// goleveldb db
	// cfg := new(config.Config)
	// cfg.DataDir = "./"
	// dbs, err := nodb.Open(cfg)
	// if err != nil {
	// 	fmt.Printf("nodb: error opening db: %v", err)
	// }
	// db, _ := dbs.Select(0)
	// err = db.Set([]byte("name"), []byte("diogo"))
	// if err != nil {
	// 	fmt.Printf("nodb: error opening db: %v", err)
	// }
	// value, _ := db.Get([]byte("name"))

	// log.Debugf("run :%s", value)

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "")
	})

	e.GET("/hello", Hello)

	// 登录页
	e.GET("/login", Login)
	e.GET("/api/user", User)

	e.HTTPErrorHandler = customHTTPErrorHandler
	// Start server
	e.Logger.Fatal(e.Start(Conf.Server.Addr))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	log.Debugf("run with conf:%d", code)
	errorPage := fmt.Sprintf("public/errorPages/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

func User(c echo.Context) error {
	// log.Debugf("header %s", c.Request().Header.Get("X_host_s"))
	return c.JSON(200, map[string]interface{}{
		"title": "你好，世界",
		"uuid":  util.GetUUID(),
		"data":  "diogoxiang",
	})
}

func Hello(c echo.Context) error {
	return c.Render(200, "hello.html", map[string]interface{}{
		"title": "你好，世界",
		"uuid":  util.GetUUID(),
		"data":  "diogoxiang",
	})
}

func Login(c echo.Context) error {
	return c.Render(200, "login.html", map[string]interface{}{
		"title": "你好，世界",
		"uuid":  util.GetUUID(),
		"data":  "diogoxiadng",
	})
}
