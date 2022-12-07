### go2ban.conf
###### firewall=auto
###### log_dir=/var/log/go2ban
###### block_list_dir=
###### grpc_port=
###### rest_port=
###### white_list=192.168.0.1 192.168.0.* //TODO
###### save_days=30
###### alert_ip=1000
###### fake_socks_ports=22 3389
###### fake_socks_fails=2
###### local_service_check_minutes=2
###### local_service_fails=2
###### Checking for wrong login attempts to local services
###### Services can have any name
###### {
######  "Service":[
######    {"On":true,"Name":"postree11","Regxp": "","LogFile":"root"},
######    {"On":true,"Name":"sshd","Regxp": "","LogFile":"sys"}
######  ]
###### }
