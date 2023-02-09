[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
![Coverage](https://img.shields.io/badge/Coverage-10.9%25-red)
[![Go Reference](https://pkg.go.dev/badge/github.com/vv198x/go2ban.svg)](https://pkg.go.dev/github.com/vv198x/go2ban)
[![Go Report Card](https://goreportcard.com/badge/github.com/vv198x/go2ban)](https://goreportcard.com/report/github.com/vv198x/go2ban)

## go2ban
go2ban is a service for protecting VDS and VPS from brute-force passwords, scanners, and DDoS attacks. It uses firewall rules to block malicious IPs and also has features such as a white list, fake SOCKS ports, and a REST server for manual IP blocking.

### Installation
1. Prerequisites: Make sure that you have a working Go development environment and that you have Go version >=1.15 installed on your machine.

2. Clone the repository:
``` 
git clone https://github.com/vv198x/go2ban.git
 ```

3. Build the binary:
``` 
make
``` 

4. Run the installer:
``` 
sudo make install
```    
   
5. Configure go2ban by editing the configuration file:
``` 
vi /etc/go2ban/go2ban.conf
```    
6. Start the service:
``` 
systemctl start go2ban
```    

7. Enable the service:
``` 
systemctl enable go2ban
```    



### Configuration
The config file allows for various settings to be customized, including:

* **firewall**: can be set to "auto" for automatic firewall rule management or "off" to disable firewall functionality.
* **log_dir**: directory for go2ban log files.
* **white_list**: IP addresses or subnets that will never be blocked.
* **grpc_port**: port for gRPC communication (default is off).
* **blocked_ips**: maximum number of IPs that can be blocked at one time.
* **fake_socks_ports**: local ports to be opened and appear to the scanner as open, but will not respond to connections (default is off).
* **fake_socks_fails**: number of failed connection attempts before an IP is blocked.
* **rest_port**: port for manual IP blocking through REST requests (default is off).
* **local_service_check_minutes**: frequency of checking local services for brute force attempts.
* **local_service_fails**: number of failed attempts before an IP is blocked.
* Additionally, the config file allows for the customization of **local service** checks, with options to specify the service name, regular expression for detecting failed attempts, and log file location. You can tell docker to check all files syslog in containers.

### Command-line:
```
-cfgFile
   Path to file go2ban.conf (default "/etc/go2ban/go2ban.conf")
-clear
   Unlock all
-d
   Run as daemon
```

### Usage
go2ban runs as a background service, continually monitoring for malicious IPs and applying firewall rules as necessary. The service can be controlled through gRPC commands or by sending REST requests to the specified port.

### Blocking in the iptables(netfilter) raw table has several advantages, including:

**Speed**: The raw table is the earliest table in the iptables firewall rule evaluation, allowing for quick and efficient blocking of incoming packets.

**Security**: The raw table provides a strong first line of defense against incoming network traffic, helping to prevent malicious activity from reaching other parts of the system.

**Connection don't established**: Blocking traffic in the raw table ensures that the connection never even opens, which can be useful for mitigating DDoS attacks and reducing CPU load.

### Development
The go2ban service is developed in Go and makes use of iptables for firewall management. The codebase is open-source and contributions are welcome.

### Changelog
For a detailed list of changes made in each version, please refer to the change.log file in the repository.

### Support
If you encounter any issues or have any questions, please open an issue in the repository or contact the developer.