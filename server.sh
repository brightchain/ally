!#/bin/bash

ps -ef | grep ally | grep -v grep | awk '{print $2}' | xargs kill -9

nohup ./ally &