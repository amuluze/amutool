#!/bin/bash
#################################
# Author     : Amu
# Date       : ${DATE} ${TIME}
# Description: 创建 topic 默认一个分区一个副本
#################################

docker exec -i kafka /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic $1 --replication-factor 1 --partitions 5
