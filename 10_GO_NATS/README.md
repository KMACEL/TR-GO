# NATS 

# Sunucuya Nats Server Kurulumu

go get github.com/nats-io/nats-server

sudo mkdir -p /srv/nats/bin
sudo mv go/bin/nats-server /srv/nats/bin

sudo nano /srv/nats/nats-server.config
    port: 4222
    net: '127.0.0.1'

Güvenlik İçin
sudo adduser --system --group --no-create-home --shell /bin/false nats
sudo chown -R nats:nats /srv

Servis Oluşturma
sudo nano /etc/systemd/system/nats.service
```bash
[Unit]
Description=NATS messaging server

[Service]
ExecStart=/srv/nats/bin/nats-server -m 8222  -c /srv/nats/nats-server.config
User=nats
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

sudo systemctl start nats
sudo systemctl enable nats

Seritifika oluşturma
mkdir ~/priv
openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout priv/nats-server.key -out priv/nats-server.crt -subj "/C=TR/ST=Turkey/L=Istanbul/O=ASPAR/CN=mke.systems"

sudo mv ~/priv /srv/nats

    sudo chmod 440 /srv/nats/priv/*
    sudo chmod 550 /srv/nats/priv
    sudo chown -R nats:nats /srv/nats/priv


sudo nano /srv/nats/gnatsd.config
tls {
  cert_file: "/srv/nats/priv/nats-server.crt"
  key_file: "/srv/nats/priv/nats-server.key"
  timeout: 1
}


Kullanıcı eklemek
authorization {
  user: user1
  password: pass1
}

sudo chown nats /srv/nats/nats-server.config 
sudo chmod 400 /srv/nats/nats-server.config 

dünyaya servisi açmak
sudo nano /srv/nats/gnatsd.config
net: '0.0.0.0'
# Kaynak

> https://nats.io/

> https://github.com/nats-io/gnatsd

> https://github.com/nats-io/go-nats

> https://github.com/nats-io/go-nats-examples

> https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-nats-on-ubuntu-16-04