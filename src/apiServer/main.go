package main

import (
	"html/template"
	"io"
	// "log"
	"net/http"

	. "apiServer/conf"
	"apiServer/module/log"

	"github.com/echo-contrib/pongor"
	"github.com/hb-go/echo-mw/captcha"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
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

	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("public/views/*.html")),
	// }

	r := pongor.GetRenderer(pongor.PongorOption{
		// 默认模版目录
		Directory: "templates",
		// if you want to reload template every request, set Reload to true.
		Reload: true,
	})
	// fmt.Println(c.NewLen(6))
	e.Renderer = r
	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, Conf.App.Name)
	})

	e.GET("/hello", Hello)
	// log.Global.Printf("ddd")

	// Start server
	e.Logger.Fatal(e.Start(Conf.Server.Addr))
}

func Hello(c echo.Context) error {
	return c.Render(200, "index.html", map[string]interface{}{
		"title": "你好，世界",
	})
}