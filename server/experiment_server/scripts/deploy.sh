#!/bin/bash
f1() {
    scp experiment_server.tar.gz musk:/tmp/
    scp experiment_server.tar.gz ahan:/tmp/
    rm -rf /data/psychology/experiment_server/*
    cp -r dist/* /data/psychology/experiment_server/
    cd /data/psychology/experiment_server/ && sh restart.sh &
}

f2() {
    ssh ahan "rm -rf /data/psychology/experiment_server/* && cd /tmp/ && tar -xvzf experiment_server.tar.gz && cp -r dist/* /data/psychology/experiment_server/ && cd /data/psychology/experiment_server/ && chmod 777 * && sh restart.sh"
}

f3() {
    ssh musk "rm -rf /data/psychology/experiment_server/* && cd /tmp/ && tar -xvzf experiment_server.tar.gz && cp -r dist/* /data/psychology/experiment_server/ && cd /data/psychology/experiment_server/ && chmod 777 * &&  sh restart.sh"
}

deploy() {
    tar -cvzf experiment_server.tar.gz dist/
    f1
    f2
    f3
    rm -rf experiment_server.tar.gz
}

deploy