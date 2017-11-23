# web

This is web system

# 开发记录

* apiServer

## Http2 生成本地 TLS 证书

go run c:\Go\src\crypto\tls\generate_cert.go --host localhost

### golang.org 设置 mirror

* $ glide mirror set https://golang.org/x/mobile
  https://github.com/golang/mobile --vcs git
* $ glide mirror set https://golang.org/x/crypto
  https://github.com/golang/crypto --vcs git
* $ glide mirror set https://golang.org/x/net https://github.com/golang/net
  --vcs git
* $ glide mirror set https://golang.org/x/tools https://github.com/golang/tools
  --vcs git
* $ glide mirror set https://golang.org/x/text https://github.com/golang/text
  --vcs git
* $ glide mirror set https://golang.org/x/image https://github.com/golang/image
  --vcs git
* $ glide mirror set https://golang.org/x/sys https://github.com/golang/sys
  --vcs git
* $ glide mirror set https://golang.org/x/sys/unix
  https://github.com/golang/sys/unix --vcs git
* $ glide mirror set golang.org/x/crypto github.com/golang/crypto --vcs git

---

vs code plugin Update( 需要重新更新的 vs 插件 )

* glide get -u -v github.com/uudashr/gopkgs
* glide get -u -v github.com/ramya-rao-a/go-outline
