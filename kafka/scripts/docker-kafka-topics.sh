#!/bin/bash
#################################
# Author     : Amu
# Date       : ${DATE} ${TIME}
# Description: 重启服务
#################################

docker exec -i kafka /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list