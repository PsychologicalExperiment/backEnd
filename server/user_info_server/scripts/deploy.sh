#!/bin/bash
f1() {
    scp user_info_server.tar.gz musk:/tmp/
    scp user_info_server.tar.gz ahan:/tmp/
    rm -rf /data/psychology/user_info_server/*
    cp -r dist/* /data/psychology/user_info_server/
    cd /data/psychology/user_info_server/ && sh restart.sh &
}

f2() {
    ssh ahan "rm -rf /data/psychology/user_info_server/* && cd /tmp/ && tar -xvzf user_info_server.tar.gz && cp -r dist/* /data/psychology/user_info_server/ && cd /data/psychology/user_info_server/ && chmod 777 * && sh restart.sh"
}

f3() {
    ssh musk "rm -rf /data/psychology/user_info_server/* && cd /tmp/ && tar -xvzf user_info_server.tar.gz && cp -r dist/* /data/psychology/user_info_server/ && cd /data/psychology/user_info_server/ && chmod 777 * &&  sh restart.sh"
}

deploy() {
    tar -cvzf user_info_server.tar.gz dist/
    f1
    f2
    f3
    rm -rf user_info_server.tar.gz
}

deploy