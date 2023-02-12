### Proxy
A simple HTTPS proxy with support for basic authorization and IP hiding. It also includes a feature to block IPs after a certain number of unsuccessful login attempts.
* IP blocking after unsuccessful login attempts
* Basic Authorization
* Hiding IP Address

### Usage
To use the proxy, you can use the curl command with the -x flag to specify the proxy address and credentials, like so:

```
curl -vv -x u:p@192.168.122.161:51461 -L https://2ip.ru/
```

### Building
To build the proxy, you can use the following command:

```
go build -o proxy
```

### Docker
You can also run the proxy using Docker and Docker Compose. To do this, you can use the following command:

```
docker-compose up
```
### Configuration
You can configure the proxy using command-line flags:
```
-addr string
proxy address (default ":51461")

-go2ban string
go2ban gRPC address (default "1.1.1.1:2048")

-pass string
Auth password (default "pass")

-user string
Auth user name (default "user")
```