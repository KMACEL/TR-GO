# MongoDB

## Kurulum

MongoDB'nin Projelerimizde MongoDB'inin Community versiyonunu kullanacağız.

### MongoDB Kurulum

> <https://docs.mongodb.com/manual/administration/install-community/>

İlk olarak paket yöneticisine MongoDB'nin imzasını ekliyoruz;

```bash
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4
```

Ardından kaynağı ekliyoruz;

```bash
echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.0.list
```

Değişikliklerimizin geçerli olması için paket yöneticisini güncelliyoruz;

```bash
sudo apt-get update
```

Son sürümün yüklenme işlemini gerçekleştiriyoruz;

```bash
sudo apt-get install -y mongodb-org
```

İşlem bittiğinde servisi çalıştırmak için aşağıdaki komutu kullanıyoruz;

```bash
sudo service mongod start
```

MongoDB'yi durdurmak için;

```bash
sudo service mongod stop
```

MongoDB'yi yeniden başlatmak için;

```bash
sudo service mongod restart
```

MongoDB'yi konsoldan çalıştırmak için;

```bash
mongo
```

MongoDB'yi bilgisayarınızdan kaldırmak için;

```bash
sudo service mongod stop
sudo apt-get purge mongodb-org*
sudo rm -r /var/log/mongodb
sudo rm -r /var/lib/mongodb
```

### GO Driver Kurulumu

GO'da paket bağımlılıklarını kontrol eden **dep** scripti bulunmaktadır. MongoDB'de, **dep** desteği bulunmaktadır. İlk önce **dep**'i yükleyelim.

```bash
go get -u github.com/golang/dep/cmd/dep
```

MongoDB, GO Driver'ı nı sisteme kurmak için;

```bash
go get -v -u go.mongodb.org/mongo-driver
cd $GOPATH/src/go.mongodb.org/mongo-driver/
git checkout -b v1.0.1 v1.0.1
dep ensure -update
dep status
```

**dep**, bir proje dosyası içine bütün bağlılıklarını ekleyen bir sistemi bulunmaktadır. Biz, örnek kullandığımız MongoDB klasörü altına aşağıdaki komutları çalıştırarak, MongoDB eklentilerini kurabiliriz;

```bash
dep init
dep ensure -add "go.mongodb.org/mongo-driver/mongo@~1.0.1"
dep ensure -update -no-vendor -v
```

**"vendor"** klasörünün altına baktığınızda **"mongodb"** bağımlılıkları eklediğini göreceksiniz. Eğer **"dep ensure -update"** derseniz ve MongoDB driverını kullanmıyorsanız, **vendor** altındaki bağımlılıkları otomatik olarak siler. ("GO" da bir şey kullanmazsanız, silinir :))

## Not

Ubuntu 16.04'teki glibc paketinin eski sürümlerinde bulunan bir hata nedeniyle, MongoDB'yi çalıştırmadan önce glibc paketini en azından glibc 2.23-0ubuntu5'e yükseltmelisiniz. Glibc paketinin eski sürümlerine sahip sistemler rasgele bellek bozulması nedeniyle veritabanı sunucusu çökmeleri ve hatalı davranışlarla karşılaşacaklardır ve MongoDB'nin üretim dağıtımları için uygun değildir

Kontrol için;

```bash
ldd --version
```

## NOT 2

Mongodb paketleri, resmi olarak MongoDB Inc. tarafından kontrol edilmekte ve geliştirilmektedir. En son MongoDB sürümleri ile güncel tutulmaktadır. Bu kurulum prosedürü, mongodb paketini kullanır.

Ubuntu tarafından sağlanan mongodb paketi, MongoDB Inc. tarafından sağlan**ma**maktadır ve mongodb işletme paketi ile çelişmektedir. Ubuntu’nun mongodb paketinin sistemde yüklü olup olmadığını kontrol etmek için **```bash sudo apt list - runed | grep mongodb```** komutunu çalıştırınız. Eğer yüklemeden önce bilgisyaarınızda kurulu bir mongodb var ise önce mongodb paketini silmelisiniz. Bunun için **```bashsudo apt remove mongodb```** ve **```bash sudo apt purge mongodb```** komutlarını kullanabilirsiniz.

## Kaynak

> <https://github.com/mongodb/mongo-go-driver>
>
> <https://docs.mongodb.com/ecosystem/drivers/go/>
>
> <https://docs.mongodb.com/manual/tutorial/getting-started/>
>
> <http://s3.amazonaws.com/info-mongodb-com/MongoDB_Enterprise_Advanced_Datasheet.pdf>
>
> <https://groups.google.com/forum/embed/#!topic/mongodb-user/sZhPCyywE-Y>
>
> <https://www.percona.com/software/mongo-database/feature-comparison>
>
> <https://golang.github.io/dep/>
>
> <https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial-part-1-connecting-using-bson-and-crud-operations>