#!/bin/bash
#################################
# Author     : Amu
# Date       : ${DATE} ${TIME}
# Description: 查看 topic 详情
#################################

docker exec -i kafka /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --describe  --topic $1
