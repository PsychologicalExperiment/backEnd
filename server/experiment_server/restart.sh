ps -ef | grep experiment_server | grep -v grep | awk  '{print $2}'| xargs kill -9 &> /dev/null 2>&1 &
nohup ./experiment_server &