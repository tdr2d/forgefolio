#!/bin/bash

set -euo pipefail

cd /tmp
curl -Lo redis.tar.gz https://download.redis.io/releases/redis-6.0.9.tar.gz
tar xzf redis.tar.gz
cd redis-6.0.9
make
cd src
sudo cp redis-cli redis-server /usr/local/bin/
