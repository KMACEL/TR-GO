# GO



#GO Tool
go get -u -v golang.org/x/tools/...
go get -u -v golang.org/x/tools/cmd/godoc
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename

go get -u -v github.com/golang/lint/golint
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/newhook/go-symbols

Kullanımı
goimports -w main.go
gofmt -w main.go
golint main.go
gomodifytags -w -file main.go -struct STRUCT_ISMI -add-tags json
go-symbols <project_path>

Vim İçin
go get -u -v github.com/nsf/gocode 

Diğer Tool
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/lukehoban/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/cweill/gotests/...