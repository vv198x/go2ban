#!/usr/bin/sh
if [[ $EUID -ne 0 ]]; then
  echo "You must run this with superuser priviliges.  Try \"sudo ./install.sh\"" 2>&1
  exit 1
fi
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
