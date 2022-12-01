#!/usr/bin/sh
grpcurl -proto 2ban.proto \
    -d '{"ip": "Test"}' \
    127.0.0.1:1024 \
    IP2ban.IPStringRequest