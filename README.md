### go2ban.conf
``` shell
firewall=auto
#log_dir=/var/log/go2ban
#white_list=192.168.0.1 192.168.0.*
#grpc_port=2048/tcp #default off
#blocked_ips=1000
#fake_socks_ports=21 3389 #default off
#fake_socks_fails=2
#rest_port= #default off
#local_service_check_minutes=
#local_service_fails=2
#Checking for wrong login attempts to local services
#Services can have any name
{
  "Service":[
    {"On":true,"Name":"sshd","Regxp": "","LogFile":"sys"},
    {"On":true,"Name":"postree11","Regxp": "","LogFile":"root"}
  ]
}

```