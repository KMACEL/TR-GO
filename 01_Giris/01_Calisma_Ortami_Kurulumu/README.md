# Çalışma Ortamı

* ### Hangi İşletim Sistemi Kullanıyoruz
Biz geliştirmeleri **"Linux Mint 18.3 Sylvia"** işletim sisteminde gerçekleştiriyoruz. İndirmek için;

  > https://linuxmint.com/edition.php?id=246

* ### Sistem Gereksinimleri
Cihazınıza bazı indirmeniz gereken programlar bulunmaktadır. Bunları indirmek için terminal ekranınızda aşağıdaki komutları çalıştırınız.

 ``` bash
sudo apt-get install g++ -y
sudo apt-get install build-essential -y
sudo apt-get install git -y
sudo apt-get install make -y
sudo apt-get install xterm -y
```


* ### IDE Nedir ?
**IDE** kısaltmasının açılmış hali **"Integrated Development Environment"** demektir. Türkçe çevirisi **"Entegre Geliştirme Ortamı"** dır.

  Aslında **IDE** bizim **yazılım projelerimizi**, rahat **programlayabilmek** için bir çok imkan sağlayan **geliştirme ortamıdır**. **Renklendirme, otomatik düzenleme, hata denetimi gerçekleştirme, derleme, çalıştırma, fonksiyonlar ve sınıflar arası gezinme** gibi bir çok kolaylıklar sağlamaktadır.

  Proje ne kadar büyük ve karmaşıksa, **IDE** gereksinimi o kadar artmaktadır.  

* ### Hangi IDE'leri Kullanabiliriz
**GO** için bir çok **IDE** seçeneği bulunmaktadır. Bunlardan en çok kullanılanları;

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

  ![Adım1](Resim/01)

  ##### Adım 2 :

  ![Adım2](Resim/)

  ##### Adım 3 :
  ![Adım3](Resim/)

  ##### Adım 4
  ![Adım4](Resim/)
