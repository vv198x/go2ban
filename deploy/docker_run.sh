#!/usr/bin/sh
#redis-cli -h 192.168.122.111 -a PASS
#config set stop-writes-on-bgsave-error no
docker run -d \
  -h redis \
  -e REDIS_PASSWORD=redis \
  -v $PWD/redis-data:/bitnami/redis/data \
  -p 6379:6379 \
  --name redis \
  --restart always \
 bitnami/redis:latest /bin/sh -c 'redis-server --requirepass redis'