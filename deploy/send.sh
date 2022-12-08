#!/usr/bin/sh
cd ..
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o deploy/go2ban
cd deploy
echo -install
scp go2ban n:/root/go2ban/
if [ "-install" = "$1" ]; then
  ssh n mkdir /root/go2ban
  scp install.sh n:/root/go2ban/
  scp go2ban.conf n:/root/go2ban/
  scp go2ban.service n:/root/go2ban/
  ssh n /root/go2ban/install.sh
fi
