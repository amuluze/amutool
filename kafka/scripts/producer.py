#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
@Time    : 2023/11/30 14:41
@Author  : Amu
@Desc    :
"""

from kafka3 import KafkaProducer

if __name__ == '__main__':
    producer = KafkaProducer(bootstrap_servers=["localhost:9092"])
    topic = "test"
    for i in range(1000):
        msg = "hello " + str(i)
        print(msg)
        producer.send(topic, msg.encode("utf-8"))
