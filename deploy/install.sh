#!/usr/bin/sh
cd /root/go2ban
systemctl stop go2ban
mkdir /var/log/go2ban
mkdir /etc/go2ban
cp go2ban.conf /etc/go2ban
chmod +x go2ban
cp go2ban /usr/local/bin
cp go2ban.service /etc/systemd/system
systemctl daemon-reload
systemctl start go2ban
systemctl enable go2ban
