## 常用的 tool

----2018年7月29日13:44:22
glide get -u -v github.com/ramya-rao-a/go-outline
glide get -u -v github.com/acroca/go-symbols
glide get -u -v github.com/mdempsky/gocode
glide get -u -v github.com/rogpeppe/godef

glide get -u -v golang.org/x/tools/cmd/godoc
glide get -u -v golang.org/x/tools/cmd/goimports
glide get -u -v golang.org/x/tools/cmd/gorename
glide get -u -v golang.org/x/tools/cmd/guru

glide get -u -v github.com/zmb3/gogetdoc

glide get -u -v github.com/golang/lint/golint

glide get -u -v github.com/fatih/gomodifytags

glide get -u -v github.com/sqs/goreturns

glide get -u -v github.com/cweill/gotests/...

glide get -u -v github.com/josharian/impl
go get -u -v github.com/haya14busa/goplay/cmd/goplay
glide get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
go get -u -v github.com/alecthomas/gometalinter
gometalinter --install

------

在一个终端中运行以下命令来安装 `errcheck`，这是许多可用的 linters（代码检测工具） 之一。

`glide get -u github.com/kisielk/errcheck`

然后，在你的代码目录中运行 `errcheck .`。