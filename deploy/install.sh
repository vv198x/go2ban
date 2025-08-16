#!/usr/bin/sh

# Определяем, нужно ли использовать sudo
if [ $EUID -eq 0 ]; then
  SUDO_CMD=""
else
  SUDO_CMD="sudo"
fi

cd /root/go2ban
$SUDO_CMD systemctl stop go2ban
$SUDO_CMD mkdir -p /var/log/go2ban
$SUDO_CMD mkdir -p /etc/go2ban
$SUDO_CMD cp go2ban.conf /etc/go2ban
chmod +x go2ban
$SUDO_CMD cp go2ban /usr/local/bin
$SUDO_CMD cp go2ban.service /etc/systemd/system
$SUDO_CMD systemctl daemon-reload
$SUDO_CMD systemctl start go2ban
$SUDO_CMD systemctl enable go2ban
