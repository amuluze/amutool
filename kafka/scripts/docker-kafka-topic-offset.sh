#!/bin/bash
#################################
# Author     : Amu
# Date       : ${DATE} ${TIME}
# Description: 重启服务
#################################

docker exec -i kafka /opt/bitnami/kafka/bin/kafka-run-class.sh  kafka.tools.GetOffsetShell --bootstrap-server localhost:9092 --topic $1 --time -1
