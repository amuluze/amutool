#!/bin/bash
#####################################################################
# Author     : wangjialong
# Date       : 2023/11/30 17:05:30
# Description: 查看特定consumer group 详情
#####################################################################

docker exec -i kafka /opt/bitnami/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group $1 --describe
