#!/bin/bash
#################################
# Author     : Amu
# Date       : ${DATE} ${TIME}
# Description: 重启服务
#################################

docker exec -i kafka find /opt/bitnami/kafka/libs/ -name \*kafka_\* | head -1 | grep -o '\kafka[^\n]*'