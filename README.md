# web

This is web system

# 开发记录

* apiServer

## Http2 生成本地 TLS 证书

go run c:\\Go\\src\\crypto\\tls\\generate_cert.go --host localhost

### golang.org 设置 mirror

* $ glide mirror set https://golang.org/x/mobile https://github.com/golang/mobile --vcs git
* $ glide mirror set https://golang.org/x/crypto https://github.com/golang/crypto --vcs git
* $ glide mirror set https://golang.org/x/net https://github.com/golang/net --vcs git
* $ glide mirror set https://golang.org/x/tools https://github.com/golang/tools --vcs git
* $ glide mirror set https://golang.org/x/text https://github.com/golang/text --vcs git
* $ glide mirror set https://golang.org/x/image https://github.com/golang/image --vcs git
* $ glide mirror set https://golang.org/x/sys https://github.com/golang/sys --vcs git
* $ glide mirror set https://golang.org/x/sys/unix  https://github.com/golang/sys/unix --vcs git

glide mirror set https://golang.org/x/crypto https://github.com/golang/crypto --vcs git

glide mirror set https://golang.org/x/net https://github.com/golang/net --vcs git

glide mirror set https://golang.org/x/sys https://github.com/golang/sys --vcs git

glide mirror set https://golang.org/x/text https://github.com/golang/text --vcs git

glide mirror set https://google.golang.org/grpc https://github.com/grpc/grpc-go --vcs git

glide mirror set https://google.golang.org/genproto https://github.com/google/go-genproto --vcs git


glide mirror set https://golang.org/x/net/websocket https://github.com/golang/net/tree/master/websocket  --vcs git

---

vs code plugin Update( 需要重新更新的 vs 插件 )

* glide get -u -v github.com/uudashr/gopkgs
* glide get -u -v github.com/ramya-rao-a/go-outline



* glide get -u github.com/766b/go-outliner

##  一些常用插件无法使用的问题
mkdir -p $GOPATH/src/golang.org/x/
cd !$
git clone https://github.com/golang/net.git
git clone https://github.com/golang/sys.git
git clone https://github.com/golang/tools.git
git clone https://github.com/golang/crypto.git

-----


glide get  -u -v github.com/uudashr/gopkgs/cmd/gopkgs



go get -u -v github.com/mdempsky/gocode 
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs 
go get -u -v github.com/ramya-rao-a/go-outline 
go get -u -v github.com/acroca/go-symbols 
go get -u -v golang.org/x/tools/cmd/guru 
go get -u -v golang.org/x/tools/cmd/gorename 
go get -u -v github.com/derekparker/delve/cmd/dlv 
go get -u -v github.com/stamblerre/gocode 
go get -u -v github.com/rogpeppe/godef 
go get -u -v github.com/ianthehat/godef 
go get -u -v github.com/sqs/goreturns 
go get -u -v golang.org/x/lint/golint 
go get -u -v github.com/cweill/gotests/... 
go get -u -v github.com/fatih/gomodifytags 
go get -u -v github.com/josharian/impl 
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct 
go get -u -v github.com/haya14busa/goplay/cmd/goplay 
go get -u -v github.com/golang/lint/golint

go get -u -v github.com/alecthomas/gometalinter -- 
gometalinter --install


go install github.com/ramya-rao-a/go-outline
go install github.com/acroca/go-symbols
go install golang.org/x/tools/cmd/guru
go install golang.org/x/tools/cmd/gorename
go install github.com/josharian/impl
go install github.com/rogpeppe/godef
go install github.com/sqs/goreturns
go install golang.org/x/lint/golint
go install github.com/cweill/gotests/gotests

go install github.com/sqs/goreturns


go install github.com/uudashr/gopkgs/cmd/gopkgs