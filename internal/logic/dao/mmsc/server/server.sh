#!/bin/sh

set -e

NSQ_SERVER_PATH="./build"

# 标准输出和错误输出均送往/dev/null
# nohup $NSQ_SERVER_PATH/nsqlookupd  1>/dev/null 2>&1 &
nohup $NSQ_SERVER_PATH/nsqd --lookupd-tcp-address=127.0.0.1:4160 1>/dev/null 2>&1 &
# nohup $NSQ_SERVER_PATH/nsqadmin --lookupd-http-address=127.0.0.1:4161 1>/dev/null 2>&1 &
