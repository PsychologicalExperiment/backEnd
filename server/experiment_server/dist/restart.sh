ps -ef | grep experiment_server | grep -v grep | grep -v ssh | grep -v scp | awk  '{print $2}'| xargs kill -9
nohup ./experiment_server &