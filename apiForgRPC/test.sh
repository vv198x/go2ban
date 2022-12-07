#!/usr/bin/sh
grpcurl -plaintext -proto 2ban.proto \
    -d '{"ip": "111.222.111.222:1"}' \
    127.0.0.1:2048\
    IP2ban.IP