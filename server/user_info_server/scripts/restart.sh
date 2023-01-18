#!/bin/env sh
ps -ef | grep user_info_server | grep -v grep | grep -v ssh | grep -v scp | awk  '{print $2}'| xargs kill -9 || true 
nohup ./user_info_server -conf=/data/psychology/user_info_server/config/config.yaml &> /dev/null 2>&1 &