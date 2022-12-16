#!/usr/bin/sh
grpcurl -plaintext -proto 2ban.proto \
    -d '{"ip": "127.0.0.2:TEST"}' \
    127.0.0.1:2048\
    IP2ban.IP