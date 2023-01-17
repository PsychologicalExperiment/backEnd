#!/bin/bash
sudo rm -rf /data/psychology/experiment_server/experiment_server &> /dev/null &&
cp bin/experiment_server /data/psychology/experiment_server/experiment_server &&
cp -r config /data/psychology/experiment_server/config &&
cp restart.sh /data/psychology/experiment_server/ &&
cd /data/psychology/experiment_server && sh restart.sh &


ssh ahan "rm -rf /data/psychology/experiment_server/experiment_server &> /dev/null" &&
scp bin/experiment_server ahan:/data/psychology/experiment_server/experiment_server &&
scp -r config ahan:/data/psychology/experiment_server/config &&
scp restart.sh ahan:/data/psychology/experiment_server/ &&
ssh ahan "cd /data/psychology/experiment_server && sh restart.sh &> /dev/null 2>&1 &"


ssh musk "rm -rf /data/psychology/experiment_server/experiment_server &> /dev/null" &&
scp bin/experiment_server musk:/data/psychology/experiment_server/experiment_server &&
scp -r config musk:/data/psychology/experiment_server/config &&
scp restart.sh musk:/data/psychology/experiment_server/ &&
ssh musk "cd /data/psychology/experiment_server && sh restart.sh &> /dev/null 2>&1 &"
