# Çalışma Ortamı

* ### Hangi İşletim Sistemi Kullanıyoruz
Biz geliştirmeleri **"Ubuntu 18.04 LTS - bionic"** işletim sisteminde gerçekleştiriyoruz. İndirmek için;

  > https://www.ubuntu.com/download/desktop

* ### Sistem Gereksinimleri
Cihazınıza bazı indirmeniz gereken programlar bulunmaktadır. Bunları indirmek için terminal ekranınızda aşağıdaki komutları çalıştırınız.

* #### Git Yüklemek
 ``` bash
sudo apt-get install git -y
```
* #### GO Yüklemek
``` bash
wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz

mkdir ~/go
mkdir ~/go/pkg
mkdir ~/go/src
mkdir ~/go/bin

echo "#GOLANG Settings" >> ~/.profile
echo export GOPATH=\$HOME/go >> ~/.profile
echo export PATH=\$PATH:/usr/local/go/bin >> ~/.profile

sudo su
echo export PATH=\$PATH:/usr/local/go/bin >> /etc/profile
exit

source ~/.profile
```
Bu işlemlerden sonra cihazınızı yeniden başlatınız.

* ### IDE Nedir ?
**IDE** kısaltmasının açılmış hali **"Integrated Development Environment"** demektir. Türkçe çevirisi **"Entegre Geliştirme Ortamı"** dır.

  Aslında **IDE** bizim **yazılım projelerimizi**, rahat **programlayabilmek** için bir çok imkan sağlayan **geliştirme ortamıdır**. **Renklendirme, otomatik düzenleme, hata denetimi gerçekleştirme, derleme, çalıştırma, fonksiyonlar ve sınıflar arası gezinme** gibi bir çok kolaylıklar sağlamaktadır.

  Proje ne kadar büyük ve karmaşıksa, **IDE** gereksinimi o kadar artmaktadır.  

* ### Hangi IDE'leri Kullanabiliriz
**GO** için bir çok **IDE** seçeneği bulunmaktadır. Bunlardan en çok kullanılanları;

## Geliştirme Ortamları
> https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

> Goland (ücretli)

  #### VIM ;
  > https://www.vim.org/

  #### Visual Code ;
  > https://code.visualstudio.com/

  #### Atom
  > https://atom.io
  


* ### Biz Hangi IDE de Kod Geliştirdik, Neden?
Biz kodlarımızı **"Visual Code"** da geliştirdk. Bunun amacı, çoğu **IDE** de olduğu gibi otomatik bazı işleri **yapmıyor** oluşu. Yani program sırasında otomatik bazı dosyaları oluşturmuyor. Yada arka planda kolaylıklar **sağlamıyor**. Her şey bizim kontrolümüz altında olıyor. **İlk öğrenme sırasında** otomatik olan işlemler, belki bazı konuları kavramada zorluk yaşatabilir. Fakat siz gönlünüzden hangi ortamda kod yazmak istiyorsanız onu seçebilirsiniz. Hatta hiç **IDE** kullanmadan, basit bir **text editöründe** dahi programlayabilirsiniz.

* ### Nasıl Yüklenir
İlk adım olarak terminalden aşağıdaki kodları çalıştırarak **"Visual Code"** editörü indirip kuruyoruz.
``` bash
wget https://go.microsoft.com/fwlink/?LinkID=760868 code_1.31.0-1549443364_amd64.deb -O code_1.31.0-1549443364_amd64.deb
sudo dpkg -i code_1.31.0-1549443364_amd64.deb
```

  **"Visual Code"** bir çok programlama dili için eklentiler sağlamaktadır. Bizde **"GO"** için bir eklenti kullanacağız.

  ##### Adım 1 :

  ![Adım1](Resim/01_VC.png)

  ##### Adım 2 :

  ![Adım2](Resim/02_VC.png)

  ##### Adım 3 :
  ![Adım3](Resim/03_VC.png)

  ##### Adım 4
  ![Adım4](Resim/04_VC.png)

  ##### Adım 5
  ![Adım4](Resim/05_VC.png)

### VIM İşlemleri

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
## .vimrc :

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
```

### Otomatik Tamamlamak İçin
```bash
cd ~/.vim/bundle
git clone https://github.com/Shougo/neocomplete.vim.git
```

#### .vimrc :

```bash
let g:neocomplete#enable_at_startup = 1
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

## Alternatif GO Tool
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