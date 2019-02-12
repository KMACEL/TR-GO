# GO

## GO Kurulum

## GO Tool
```bash
go get -u -v golang.org/x/tools/...
go get -u -v golang.org/x/tools/cmd/godoc
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename

go get -u -v github.com/golang/lint/golint
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/newhook/go-symbols
```

### Kullanımı
> goimports -w main.go
> gofmt -w main.go
> golint main.go
> gomodifytags -w -file main.go -struct STRUCT_ISMI -add-tags json
> go-symbols <project_path>


## Diğer Tool
```bash
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/lukehoban/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/cweill/gotests/...
```

## Geliştirme Ortamları
> https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

> VIM

> Visual Code

> ATOM

> Goland (ücretli)

## VIM için

> http://blog.oguzhan.info/?p=1248

```bash
sudo apt-get install git curl wget build-essential vim neovim vim-nox vim-pathogen
git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go
git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go
mkdir .vim/bundle

mkdir -p ~/.vim/autoload ~/.vim/bundle
curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim
```


### Otomatik Tamamlamak İçin
```bash
cd ~/.vim/bundle
git clone https://github.com/Shougo/neocomplete.vim.git
```

### Hata Bulucu

```bash
cd ~/.vim/bundle
git clone https://github.com/scrooloose/syntastic
```
#### Kullanımı
> :SyntasticCheck go

#### .vimrc Eklentisi :
```bash
set statusline+=%#warningmsg#
set statusline+=%{SyntasticStatuslineFlag()}
set statusline+=%*

let g:syntastic_always_populate_loc_list = 1
let g:syntastic_auto_loc_list = 1
let g:syntastic_check_on_open = 1
let g:syntastic_check_on_wq = 0
```


## Genel .vimrc :

```bash
call pathogen#infect()

filetype plugin on

" Highlight
let g:go_disable_autoinstall = 0
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_structs = 1
let g:go_highlight_operators = 1
let g:go_highlight_build_constraints = 1

let g:neocomplete#enable_at_startup = 1
```

## GO Eklentilerini İndirmek İçin :

> vim in :GoInstallBinaries


## Vim Komutları

> :Explore - Dosya Gezgini
> :NERDTreeToggl - Dosy Gezgini



### İsteğe Bağlı TagBAR
```bash 
go get -u github.com/jstemmer/gotags
cd ~/.vim/bundle
git clone https://github.com/majutsushi/tagbar.git
```

#### .vimrc Eklentisi :

```bash
let g:tagbar_type_go = {  
    \ 'ctagstype' : 'go',
    \ 'kinds'     : [
        \ 'p:package',
        \ 'i:imports:1',
        \ 'c:constants',
        \ 'v:variables',
        \ 't:types',
        \ 'n:interfaces',
        \ 'w:fields',
        \ 'e:embedded',
        \ 'm:methods',
        \ 'r:constructor',
        \ 'f:functions'
    \ ],
    \ 'sro' : '.',
    \ 'kind2scope' : {
        \ 't' : 'ctype',
        \ 'n' : 'ntype'
    \ },
    \ 'scope2kind' : {
        \ 'ctype' : 't',
        \ 'ntype' : 'n'
    \ },
    \ 'ctagsbin'  : 'gotags',
    \ 'ctagsargs' : '-sort -silent'
\ }
```